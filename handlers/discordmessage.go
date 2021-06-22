package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/commands"
	log "github.com/sirupsen/logrus"
)

func DiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Filter messages with less than 2 characters
	if len(m.Message.Content) >= 1 {

		// Message started with !, probably a command.
		// No reason to act on commands without it.
		if (m.Message.Content[:1]) == "!" {

			// Split the first character off the message
			commandName := strings.SplitN(m.Message.Content, " ", 2)[0][1:]
			log.Println("commandname:", commandName)

			cmdParams := strings.SplitN(m.Message.Content[1:], " ", 500)
			log.Println("cmdParams:", cmdParams)

			// Handle how many characters the message contains.
			msgLen := len(strings.SplitN(m.Message.Content, " ", -1))
			log.Println("msgLen: ", msgLen)

			// log.Info(m.Message)
			if m.Message.Author.ID == s.State.User.ID {
				return
			}
			if m.Content == "!ping" {
				s.ChannelMessageSend(m.ChannelID, "Pong!")
			}
			if m.Content == "!pong" {
				s.ChannelMessageSend(m.ChannelID, "Ping!")
			}
			if m.Content == "!test" {
				commands.Test(s, m)
			}

		}
	}
}
