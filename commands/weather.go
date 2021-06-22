package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Weather(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Message.Content) == "!weather" {
		// No location was provided
		s.ChannelMessageSend(m.ChannelID, "Usage: !weather [location]")
	} else {
		// Strip command name from the message and
		// assume the rest of the message is the location.
		location := m.Message.Content[9:]
		log.Info(location)
		resp, err := http.Get(fmt.Sprintf("https://customapi.aidenwallis.co.uk/api/v1/misc/weather/%s", location))
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
}
