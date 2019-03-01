package slack

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func SendMessage(message string, channel string) {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("No token set")
	}
	api := slack.New(token)
	p := slack.PostMessageParameters{}

	_, _, err := api.PostMessage(channel, message, p)
	if err != nil { //TODO: propagate err
		fmt.Println(err)
	}
}
