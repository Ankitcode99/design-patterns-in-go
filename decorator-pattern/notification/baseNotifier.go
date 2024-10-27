package notification

import "fmt"

type BasicNotifier struct {
	channel string
}

func NewBasicNotifier(channel string) *BasicNotifier {
	return &BasicNotifier{channel: channel}
}

func (n *BasicNotifier) Notify(message string) error {
	fmt.Printf("\nSending basic notification via %s: %s\n", n.channel, message)
	return nil
}
