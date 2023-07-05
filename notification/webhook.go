package notification

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

var WEBHOOK_URL string = ""

func Alert(content string) {
	var username = "Node Alert"

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(WEBHOOK_URL, message)
	if err != nil {
		log.Fatal(err)
	}
}

func SetWebhook(url string) {
	WEBHOOK_URL = url
}

func GetWekhook() string {
	return WEBHOOK_URL
}
