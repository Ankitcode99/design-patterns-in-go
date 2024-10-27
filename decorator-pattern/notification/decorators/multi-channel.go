package decorators

import notification "github.com/Ankitcode99/design-patterns-in-go/decorator-pattern/notification"

type MultiChannelDecorator struct {
	BaseDecorator
	additionalChannels []notification.Notifier
}

func NewMultiChannelDecorator(n notification.Notifier, additionalChannels []notification.Notifier) *MultiChannelDecorator {
	return &MultiChannelDecorator{
		BaseDecorator:      BaseDecorator{wrapped: n},
		additionalChannels: additionalChannels,
	}
}

func (d *MultiChannelDecorator) Notify(message string) error {
	err := d.wrapped.Notify(message)
	if err != nil {
		return err
	}

	for _, channel := range d.additionalChannels {
		if err := channel.Notify(message); err != nil {
			return err
		}
	}
	return nil
}
