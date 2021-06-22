package commands

import (
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Eightball(s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://customapi.aidenwallis.co.uk/api/v1/misc/8ball")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong :(")
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong :(")
		log.Error(err)
	}

	s.ChannelMessageSend(m.ChannelID, string(body))
}
