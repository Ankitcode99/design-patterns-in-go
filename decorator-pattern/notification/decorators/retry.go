package decorators

import (
	"fmt"
	"time"

	notification "github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification"
)

type RetryDecorator struct {
	BaseDecorator
	maxRetries int
}

func NewRetryDecorator(n notification.Notifier, maxRetries int) *RetryDecorator { // Changed parameter type
	return &RetryDecorator{
		BaseDecorator: BaseDecorator{wrapped: n},
		maxRetries:    maxRetries,
	}
}

func (d *RetryDecorator) Notify(message string) error {
	var err error
	for i := 0; i < d.maxRetries; i++ {
		err = d.wrapped.Notify(message)
		if err == nil {
			fmt.Printf("\n\n sent notification in %d attempt! \n\n", i+1)
			return nil
		}
		fmt.Printf("Retry attempt %d/%d failed: %v\n", i+1, d.maxRetries, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	return fmt.Errorf("failed after %d retries: %v", d.maxRetries, err)
}
