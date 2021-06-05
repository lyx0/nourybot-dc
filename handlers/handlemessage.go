package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/commands"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Don't act on  bots own messages
	if m.Message.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ping" {
		commands.HandlePing(s, m)
	}
	if m.Content == "!pong" {
		commands.HandlePong(s, m)
	}
}
