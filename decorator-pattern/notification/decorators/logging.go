package decorators

import (
	"fmt"
	"time"

	notification "github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification"
)

type LoggingDecorator struct {
	BaseDecorator
}

func NewLoggingDecorator(n notification.Notifier) *LoggingDecorator { // Changed parameter type
	return &LoggingDecorator{BaseDecorator{wrapped: n}}
}

func (d *LoggingDecorator) Notify(message string) error {
	fmt.Printf("\n\n[%s] Sending notification...\n", time.Now().Format(time.RFC3339))
	err := d.wrapped.Notify(message)
	if err != nil {
		fmt.Printf("[%s] Error sending notification: %v\n", time.Now().Format(time.RFC3339), err)
	} else {
		fmt.Printf("[%s] Notification sent successfully\n", time.Now().Format(time.RFC3339))
	}
	return err
}
