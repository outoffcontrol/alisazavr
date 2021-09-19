package yandex

import (
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
)

func GetToken(l *zap.SugaredLogger, client *http.Client, login string, passwd string) string {

	// get cookie (yandexid)
	resp, err := client.Get("https://passport.yandex.ru/")
	if err != nil {
		l.Fatal("Cannot get cookie.", err.Error())
	}

	// auth with cookie and get SessionId
	resp, err = client.PostForm("https://passport.yandex.ru/passport?mode=auth&retpath=https://yandex.ru", url.Values{
        "login": {login},
        "passwd": {passwd},
    })
	if err != nil {
		l.Fatal("Cannot auth at Yandex.", err.Error())
	}

	// get csrf token
	resp, err = client.Get("https://frontend.vh.yandex.ru/csrf_token")
	if err != nil {
		l.Fatal("Cannot get csrf token.", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			l.Fatal(err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Fatal(err.Error())
	}
	if string(body) == "Can't get token" {
		l.Fatal("Can't get yandex token: ", resp.StatusCode)
	}

	return string(body)
}