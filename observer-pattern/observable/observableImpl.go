package observable

import (
	"fmt"

	models "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/models"
	observers "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/observer"
)

type Category struct {
	observers map[observers.Observer]bool
	name      string
	newsFeeds []models.NewsFeed
}

// NewCategory creates a new Category instance.
func NewCategory(name string) *Category {
	return &Category{
		name:      name,
		observers: make(map[observers.Observer]bool),
		newsFeeds: []models.NewsFeed{},
	}
}

// subscribe registers an observer to receive notifications for this category.
func (c *Category) Subscribe(observer observers.Observer) {
	c.observers[observer] = true
}

// unsubscribe removes an observer from receiving notifications for this category.
func (c *Category) Unsubscribe(observer observers.Observer) {
	delete(c.observers, observer)
}

// notifySubscribers sends notifications to all subscribed observers about a new news feed.
func (c *Category) notifySubscribers(newsFeed models.NewsFeed) {
	fmt.Printf("\nNotifying subscribers of %s category\n", c.name)
	for o := range c.observers {
		o.Notify(newsFeed)
	}
	fmt.Printf("\n\n===========================\n\n")
}

// addNewNewsFeed adds a new news feed to the category and notifies subscribed observers.
func (c *Category) AddNewNewsFeed(newsFeed models.NewsFeed) {
	c.newsFeeds = append(c.newsFeeds, newsFeed)
	c.notifySubscribers(newsFeed)
}
