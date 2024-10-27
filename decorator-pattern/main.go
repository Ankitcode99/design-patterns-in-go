package main

import (
	"fmt"

	"github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification"
	decorators "github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification/decorators"
)

func main() {
	emailNotifier := notification.NewBasicNotifier("email")

	// Create additional channels for multi-channel notification
	smsNotifier := notification.NewBasicNotifier("SMS")
	slackNotifier := notification.NewBasicNotifier("Slack")
	// Stack decorators
	notifier := decorators.NewMultiChannelDecorator(
		decorators.NewRetryDecorator(
			decorators.NewLoggingDecorator(emailNotifier),
			3,
		),
		[]notification.Notifier{smsNotifier, slackNotifier},
	)

	// Send a notification
	message := "Important system update: Server maintenance scheduled"
	err := notifier.Notify(message)
	if err != nil {
		fmt.Printf("Failed to send notification: %v\n", err)
	}
}
