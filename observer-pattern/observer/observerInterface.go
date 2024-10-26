package observer

import (
	models "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/models"
)

// Observer defines the interface for receiving notifications about new news feeds.
type Observer interface {
	Notify(models.NewsFeed)
}
