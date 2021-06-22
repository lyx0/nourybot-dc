package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/commands"
	log "github.com/sirupsen/logrus"
)

func DiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Test string
	// log.Info(m.Message)

	// Do not act on bots own messages.
	if m.Message.Author.ID == s.State.User.ID {
		return
	}

	// Filter messages with less than 1 characters.
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

			msgLen := len(strings.SplitN(m.Message.Content, " ", -2))
			log.Println("msgLen: ", msgLen)

			// Finally check if the message contains a known
			// command and act accordingly
			switch commandName {
			case "!":
				s.ChannelMessageSend(m.ChannelID, "!")

			case "8ball":
				commands.Eightball(s, m)
			case "coinflip":
				commands.Coinflip(s, m)
			case "cf":
				commands.Coinflip(s, m)
			case "coin":
				commands.Coinflip(s, m)
			case "ping":
				s.ChannelMessageSend(m.ChannelID, "Pong!")
			case "pong":
				s.ChannelMessageSend(m.ChannelID, "Ping!")
			case "test":
				commands.Test(s, m)
			case "num":
				if msgLen == 1 {
					commands.RandomNumberFact(s, m)
				} else if msgLen == 2 {
					commands.NumberFact(s, m, cmdParams[1])
				}
			case "numfact":
				if msgLen == 1 {
					commands.RandomNumberFact(s, m)
				} else if msgLen == 2 {
					commands.NumberFact(s, m, cmdParams[1])
				}
			case "number":
				if msgLen == 1 {
					commands.RandomNumberFact(s, m)
				} else if msgLen == 2 {
					commands.NumberFact(s, m, cmdParams[1])
				}
			case "numberfact":
				if msgLen == 1 {
					commands.RandomNumberFact(s, m)
				} else if msgLen == 2 {
					commands.NumberFact(s, m, cmdParams[1])
				}
			case "rxkcd":
				commands.RandomXkcd(s, m)
			case "xkcd":
				commands.Xkcd(s, m)
			case "weather":
				commands.Weather(s, m)
			}
		}
	}
}
