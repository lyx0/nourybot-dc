package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/config"
	"github.com/lyx0/nourybot-dc/handlers"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	cfg           *config.Config
	discordClient *discordgo.Session
}

func NewBot(cfg *config.Config, discordClient *discordgo.Session) *Bot {
	return &Bot{
		cfg:           cfg,
		discordClient: discordClient,
	}
}

func (b *Bot) newDiscordClient() *discordgo.Session {
	discordClient, err := discordgo.New("Bot " + b.cfg.DC_AUTH)
	if err != nil {
		log.Fatal("Error authenticating with discord", err)
	}

	return discordClient
}

func (b *Bot) ConnectDiscord() error {
	discordClient := b.newDiscordClient()

	discordClient.AddHandler(b.discordMessage)
	discordClient.Identify.Intents = discordgo.IntentsGuildMessages

	err := discordClient.Open()
	if err != nil {
		log.Fatal("Error connecting to Discord: ", err)
	}

	log.Info("Connected to Discord.")
	return err
}

func (b *Bot) discordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	handlers.DiscordMessage(s, m)
}

func (b *Bot) CloseConnection() {
	b.discordClient.Close()
}
