package telegram

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Updates struct {
	Updates []Update `json:"result"`
}

type Update struct {
	Update_id int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Message_id int `json:"message_id"`
	From From `json:"from"`
	Chat Chat `json:"chat"`
	Date int `json:"date"`
	Text string `json:"text"`
}

type From struct {
	Id int `json:"id"`
	Is_bot bool `json:"is_bot"`
	First_name string `json:"first_name"`
	Username string `json:"username"`
	Language_code string `json:"language_code"`
}

type Chat struct {
	Id int `json:"id"`
	First_name string `json:"first_name"`
	Username string `json:"username"`
	Type string `json:"type"`
}


func GetLastTelegramPrivateMessage(l *zap.SugaredLogger, client *http.Client, botUrl string) (int, string) {

	updates := Updates{}

	resp, err := client.Get(botUrl + "/getupdates?offset=-1")
	if err != nil {
		l.Error(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Error(err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &updates)
	if err != nil {
		l.Fatal(err.Error())
	}

	var updateId int
	var data string

	if len(updates.Updates) > 0 {
		updateId = updates.Updates[0].Update_id
		data = updates.Updates[0].Message.Text
	}

	return updateId, data
}