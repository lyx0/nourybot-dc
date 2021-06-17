package handlers

import "github.com/bwmarrin/discordgo"

func DiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
