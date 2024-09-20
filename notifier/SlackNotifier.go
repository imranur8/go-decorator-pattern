package notifier

import "fmt"

type SlackNotifier struct {
	BaseNotifier
}

func (s *SlackNotifier) SendSlackMessage(channel string, message string) error {
	// Simulate sending a Slack message
	fmt.Printf("Sending Slack Message to channel: %s, Message: %s\n", channel, message)
	return nil
}
