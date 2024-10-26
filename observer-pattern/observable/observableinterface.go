package observable

import (
	models "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/models"
	observers "github.com/Ankitcode99/design-patterns-in-go/observer-pattern/observer"
)

type CategoryObservable interface {
	Subscribe(observers.Observer)

	Unsubscribe(observers.Observer)

	notifySubscribers(models.NewsFeed)

	AddNewNewsFeed(models.NewsFeed)
}
