package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var SlackClient *slack.Client

func InitiateSlack() {
	slackToken := os.Getenv("SLACK_API_TOKEN") // Replace with your bot token
	SlackClient = slack.New(slackToken)
}

func SendDirectMessage(userID, messageText string) {

	// Step 1: Open a DM channel with the user
	channel, _, _, err := SlackClient.OpenConversation(&slack.OpenConversationParameters{
		Users: []string{userID},
	})
	if err != nil {
		fmt.Printf("failed to open conversation: %v", err)
		return
	}

	// Step 2: Send a message to the DM channel
	_, _, err = SlackClient.PostMessage(channel.ID, slack.MsgOptionText(messageText, false))
	if err != nil {
		fmt.Printf("failed to send message: %v", err)
		return
	}
}
