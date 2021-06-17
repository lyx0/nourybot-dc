package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-dc/bot"
	"github.com/lyx0/nourybot-dc/config"
	"github.com/lyx0/nourybot-dc/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()

	sqlClient := db.Connect(cfg)
	defer sqlClient.Close()

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)
	discordClient, err := discordgo.New("Bot " + cfg.DC_AUTH)
	if err != nil {
		log.Fatal("Couldn't connect to Discord", err)
	}

	bot := bot.NewBot(cfg, twitchClient, discordClient, sqlClient)

	err = bot.ConnectTwitch()
	if err != nil {
		log.Fatal("Couldn't connect to Twitch", err)
	}

	err = bot.ConnectDiscord()
	if err != nil {
		log.Fatal("Couldn't connect to Discord", err)
	}

}
