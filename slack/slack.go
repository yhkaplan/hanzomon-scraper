package slack

import (
	"os"

	"github.com/nlopes/slack"
)

func SendMessage(message string, channel string) {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	msg := rtm.NewOutgoingMessage(message, channel)
	rtm.SendMessage(msg)
}
