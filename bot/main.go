package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	handlers "github.com/lyx0/nourybot-dc/handlers"
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DC_AUTH := os.Getenv("DC_AUTH")
	dcb, err := discordgo.New("Bot " + DC_AUTH)
	if err != nil {
		log.Fatal("Error authenticating with discord, is our token invalid?")
	}

	// Register the messageCreate func as a callback for messageCreate events.
	dcb.AddHandler(handlers.MessageCreate)

	dcb.Identify.Intents = discordgo.IntentsGuildMessages

	err = dcb.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	fmt.Println("Connected!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dcb.Close()
}
