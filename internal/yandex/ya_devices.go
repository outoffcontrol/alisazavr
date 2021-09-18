package yandex

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Devices struct {
	Devices []Device `json:"items"`
}

type Device struct {
	Icon			string 	`json:"icon"`
	Id				string 	`json:"id"`
	Name 			string 	`json:"name"`
	Online 			bool 	`json:"online"`
	Platform		string	`json:"platform"`
	Screen_capable	bool	`json:"screen_capable"`
	Screen_present	bool	`json:"screen_present"`
}

var yaDeviceId string

func GetDevices(l *zap.SugaredLogger, client *http.Client) string {
	resp, err := client.Get("https://quasar.yandex.ru/devices_online_stats")
	if err != nil {
		l.Fatal("Cannot get devices.", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Fatal(err.Error())
	}

	devices := Devices{}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		l.Fatal(err.Error())
	}

	for i := 0; i < len(devices.Devices); i++ {
		if devices.Devices[i].Platform == "yandexstation" {
			yaDeviceId = devices.Devices[i].Id
		}
	}
	return yaDeviceId
}

