package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lyx0/nourybot-dc/pkg/common"
	"go.uber.org/zap"
)

type xkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

func Latest() (num, title, image string) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		sugar.Error(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		sugar.Error(err)
	}

	var responseObject xkcdResponse
	json.Unmarshal(responseData, &responseObject)

	return fmt.Sprint(responseObject.Num), responseObject.SafeTitle, responseObject.Img
}

func Random() (num, title, image string) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	comicNum := fmt.Sprint(common.GenerateRandomNumber(2700))
	response, err := http.Get(fmt.Sprintf("http://xkcd.com/%v/info.0.json", comicNum))
	if err != nil {
		sugar.Error(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		sugar.Error(err)
	}

	var responseObject xkcdResponse
	json.Unmarshal(responseData, &responseObject)

	return fmt.Sprint(responseObject.Num), responseObject.SafeTitle, responseObject.Img
}

func Specific(comicNum string) (num, title, image string) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	response, err := http.Get(fmt.Sprintf("http://xkcd.com/%v/info.0.json", comicNum))
	if err != nil {
		sugar.Error(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		sugar.Error(err)
	}

	var responseObject xkcdResponse
	json.Unmarshal(responseData, &responseObject)

	return fmt.Sprint(responseObject.Num), responseObject.SafeTitle, responseObject.Img
}
