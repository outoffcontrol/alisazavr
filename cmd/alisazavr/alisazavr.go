package main

import (
	"github.com/outoffcontrol/alisazavr/internal/telegram"
	"github.com/outoffcontrol/alisazavr/internal/yandex"
	"go.uber.org/zap"
	"net/http"
	"net/http/cookiejar"
	"os"
)


func main() {

	//init logger https://github.com/uber-go/zap
	productionLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() { _ = productionLogger.Sync() }()
	sugaredLogger := productionLogger.Sugar()
	sugaredLogger.Info("Starting!")

	//init http client with cookie
	jar, err := cookiejar.New(nil)
	if err != nil { }
	
	client := &http.Client{
		Jar: jar,
	}

	// credentials for yandex.passport
	login := os.Getenv("LOGIN")
	if login == "" {
		sugaredLogger.Fatal("LOGIN is empty.")
	}
	passwd := os.Getenv("PASSWD")
	if passwd == "" {
		sugaredLogger.Fatal("PASSWD is empty.")
	}
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		sugaredLogger.Fatal("BOT_TOKEN is empty.")
	}

	botUrl := "https://api.telegram.org/bot" + botToken
	update_id_temp, _ := telegram.GetLastTelegramPrivateMessage(sugaredLogger, client, botUrl)
	token := yandex.GetToken(sugaredLogger, client, login, passwd)
	device := yandex.GetDevices(sugaredLogger, client)

	for true {
		update_id, message := telegram.GetLastTelegramPrivateMessage(sugaredLogger, client, botUrl)
		if update_id != update_id_temp {
			sugaredLogger.Info("New message in telegram bot.")
			yandex.SendYoutubeToStation(sugaredLogger, message, client, device, token)
			update_id_temp = update_id
		}
	}
}