package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	directMessage "github.com/lyx0/nourybot-dc/internal/handlers/direct_messages"
	guildMessage "github.com/lyx0/nourybot-dc/internal/handlers/guild_messages"
	slashcommands "github.com/lyx0/nourybot-dc/internal/slash_commands"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/store"
	"go.uber.org/zap"
)

type config struct {
	env          string
	discordToken string
	db           struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type Application struct {
	Session *discordgo.Session
	Ken     *ken.Ken
	Db      *sql.DB
	Log     *zap.SugaredLogger
}

var envFlag string

func init() {
	flag.StringVar(&envFlag, "env", "dev", "database connection to use: (dev/prod)")
	flag.Parse()
}

func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	err := godotenv.Load()
	if err != nil {
		sugar.Fatalw("Error loading .env file",
			"err", err)
	}

	var cfg config

	switch envFlag {
	case "dev":
		cfg.db.dsn = os.Getenv("LOCAL_DSN")
	case "prod":
		cfg.db.dsn = os.Getenv("SUPABASE_DSN")
	}

	sugar.Infow("Running in: ",
		"envFlag", envFlag,
		"cfg.env", cfg.env)

	cfg.discordToken = os.Getenv("DC_BOT_TOKEN")
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"

	db, err := openDB(cfg, sugar)
	if err != nil {
		sugar.Fatalw("Could not connect to database:",
			"err", err)
	}

	session, err := discordgo.New("Bot " + cfg.discordToken)
	if err != nil {
		sugar.Errorw("Error creating Discord session:",
			"err", err)
		return
	}
	defer session.Close()

	k, err := ken.New(session, ken.Options{
		CommandStore: store.NewDefault(),
	})
	if err != nil {
		sugar.Fatalw("Error creating Ken:",
			"err", err)
	}

	err = k.RegisterCommands(
		new(slashcommands.TestCommand),
		new(slashcommands.WeatherCommand),
		new(slashcommands.CurrencyCommand),
		new(slashcommands.XkcdCommand),
	)
	if err != nil {
		sugar.Fatalw("Error registering Ken command:",
			"err", err)
	}

	defer k.Unregister()

	app := &Application{
		Session: session,
		Ken:     k,
		Db:      db,
		Log:     sugar,
	}

	session.AddHandler(directMessage.Create)
	session.AddHandler(guildMessage.Create)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = app.Session.Open()
	if err != nil {
		sugar.Errorw("Error opening Discord websocket connection:",
			"err", err)
		return
	}

	sugar.Infow("Started successfully. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func openDB(cfg config, sugar *zap.SugaredLogger) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	sugar.Infow("Database connection:",
		"dsn", cfg.db.dsn)

	return db, nil
}
