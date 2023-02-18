package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type config struct {
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
	Log *zap.SugaredLogger
}

func main() {
	var cfg config

	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	err := godotenv.Load()
	if err != nil {
		sugar.Fatalw("Error loading .env file",
			"err", err)
	}

	cfg.discordToken = os.Getenv("DC_BOT_TOKEN")

	cfg.db.dsn = os.Getenv("DB_DSN")
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"

	dg, err := discordgo.New("Bot " + cfg.discordToken)
	if err != nil {
		sugar.Errorw("Error creating Discord session:",
			"err", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	app := &Application{
		Dgs: dg,
		Log: sugar,
	}

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		sugar.Errorw("Failed to create a DM channel",
			"err:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		sugar.Errorw("Error while sending DM message:",
			"err", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DMs in your privacy settings?",
		)
	}
}
