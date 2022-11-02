package tools

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// notify.go will send a notification to a user using Telegram

func Send2Telegram(tg_token string, chat_id string, rabbitmq_url string) (string, error) {
	// Convert the following curl in Go
	// curl -XPOST https://api.telegram.org/bot{$tg_token}/sendMessage -d "message=rabbit $rabbitmq_url changed on $NOW" -d "chat_id=$chat_id"
	tg_url := "https://api.telegram.org/bot" + tg_token + "/sendMessage"
	params := url.Values{}
	params.Add("message", "rabbit $rabbitmq_url changed on $NOW")
	params.Add("chat_id", "$chat_id")
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", tg_url, body)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	return "OK", nil

}
