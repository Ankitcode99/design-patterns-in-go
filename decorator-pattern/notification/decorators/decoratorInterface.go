package decorators

import (
	n "github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification"
)

type BaseDecorator struct {
	wrapped n.Notifier
}
