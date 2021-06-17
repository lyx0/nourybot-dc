package main

import (
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		log.Info("Starting connection to Twitch")

		err = bot.ConnectTwitch()
		if err != nil {
			log.Fatal("Couldn't connect to Twitch", err)
			os.Exit(1)
		}
	}()

	go func() {
		log.Info("Starting connection to Discord")

		err = bot.ConnectDiscord()
		if err != nil {
			log.Fatal("Couldn't connect to Discord", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c

	sig := <-c
	log.Info("Got signal:", sig)
	os.Exit(0)
}
