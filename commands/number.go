package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func RandomNumberFact(s *discordgo.Session, m *discordgo.MessageCreate) {

	response, err := http.Get("http://numbersapi.com/random/trivia")
	if err != nil {
		log.Error(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	s.ChannelMessageSend(m.ChannelID, string(responseData))
}

func NumberFact(s *discordgo.Session, m *discordgo.MessageCreate, number string) {
	response, err := http.Get(fmt.Sprint("http://numbersapi.com/" + number))
	if err != nil {
		log.Error(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}
	s.ChannelMessageSend(m.ChannelID, string(responseData))
}
