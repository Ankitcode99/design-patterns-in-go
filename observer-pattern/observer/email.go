package observer

import (
	"fmt"

	models "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/models"
)

type EmailNotifier struct {
	EmailAddress string // Make the field public with uppercase first letter
}

func (e *EmailNotifier) Notify(newsFeed models.NewsFeed) {
	fmt.Printf("Sending email notification for news: %s to \t %s\n", newsFeed.Headline, e.EmailAddress)
}
