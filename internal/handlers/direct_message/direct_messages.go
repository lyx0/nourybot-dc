package directmessage

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func Create(s *discordgo.Session, m *discordgo.MessageCreate) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		sugar.Errorw("Failed to create a DM channel",
			"err:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		sugar.Errorw("Error while sending DM message:",
			"err", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DMs in your privacy settings?",
		)
	}
}
