package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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
	dcb.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore own emssages
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Message containing "ping"
	if m.Content == "ping" {
		return
	}

	// Create a priate channel with the user who sent the message
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error creating channel: ", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	// Send message through the channel we created
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		fmt.Println("error sending DM message: ", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM.",
		)
	}
}
