package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/lyx0/nourybot-dc/utils"
)

// https://xkcd.com/json.html
type XkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

func RandomXkcd(s *discordgo.Session, m *discordgo.MessageCreate) {

	comicNum := fmt.Sprint(utils.GenerateRandomNumber(2468))
	response, err := http.Get(fmt.Sprint("http://xkcd.com/" + comicNum + "/info.0.json"))
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	str := fmt.Sprint("Random Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)
	s.ChannelMessageSend(m.ChannelID, string(str))
}

func Xkcd(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject XkcdResponse
	json.Unmarshal(responseData, &responseObject)

	str := fmt.Sprint("Current Xkcd #", responseObject.Num, " Title: ", responseObject.SafeTitle, " ", responseObject.Img)

	s.ChannelMessageSend(m.ChannelID, string(str))
}
