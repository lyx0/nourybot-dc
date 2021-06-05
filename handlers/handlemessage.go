package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/commands"
)

func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ping" {
		commands.HandlePing(s, m)
	}
	if m.Content == "!pong" {
		commands.HandlePing(s, m)
	}
}
