package yandex

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
)

type Message struct {
	Msg Msg `json:"msg"`
	Device string `json:"device"`
}

type Msg struct {
	Provider_item_id string `json:"provider_item_id"`
	Player_id string `json:"player_id"`
}

func validationUrl(l *zap.SugaredLogger, url string, pattern string) bool {
	matched, err := regexp.MatchString(pattern, url)
	if err != nil {
		l.Fatal(err.Error())
	}
	return matched
}

func SendYoutubeToStation(l *zap.SugaredLogger, message string, client *http.Client, device string, token string)  {

	var result string

	if validationUrl(l, message ,"https://m.*") {
		result = strings.ReplaceAll(message, "https://m.", "https://www.")
	} else if validationUrl(l, message ,"https://youtu.be/*") {
		result = strings.ReplaceAll(message, "https://youtu.be/", "https://www.youtube.com/watch?v=")
	} else {
		result = message
	}

	msg := Msg{Provider_item_id: result, Player_id: "youtube"}

	body := &Message{
		Msg: msg,
		Device: device,
	}
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		l.Error(err.Error())
	}

	req, err := http.NewRequest("POST", "https://yandex.ru/video/station", payloadBuf)
	if err != nil {
		l.Fatal(err.Error())
	}
	req.Header.Set("x-csrf-token", token)
	resp, err := client.Do(req)
	if err != nil {
		l.Fatal(err.Error())
	}
	l.Info("Send video to station: ", resp.StatusCode)
}

