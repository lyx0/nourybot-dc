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
	Dgs *discordgo.Session
	Db  *sql.DB
	Log *zap.SugaredLogger
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

	dg, err := discordgo.New("Bot " + cfg.discordToken)
	if err != nil {
		sugar.Errorw("Error creating Discord session:",
			"err", err)
		return
	}

	db, err := openDB(cfg, sugar)
	if err != nil {
		sugar.Fatalw("Could not connect to database:",
			"err", err)
	}

	app := &Application{
		Dgs: dg,
		Db:  db,
		Log: sugar,
	}

	dg.AddHandler(directMessage.Create)
	dg.AddHandler(guildMessage.Create)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = app.Dgs.Open()
	if err != nil {
		sugar.Errorw("Error opening Discord websocket connection:",
			"err", err)
		return
	}

	sugar.Infow("Started successfully. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
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
