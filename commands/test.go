package commands

import "github.com/bwmarrin/discordgo"

func Test(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "xd")
}
