package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/bot"
	"github.com/lyx0/nourybot-dc/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()

	discordClient, err := discordgo.New("Bot " + cfg.DC_AUTH)
	if err != nil {
		log.Fatal("Couldn't connect to Discord", err)
	}

	bot := bot.NewBot(cfg, discordClient)

	log.Info("Connecting to Discord")

	err = bot.ConnectDiscord()
	if err != nil {
		log.Fatal("Couldn't connect to Discord", err)
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-c

	sig := <-c
	log.Info("Got signal:", sig)
	bot.CloseConnection()
	os.Exit(0)
}
