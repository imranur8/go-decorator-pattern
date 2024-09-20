package notifier

import "fmt"

type EmailNotifier struct {
	BaseNotifier
}

func (e *EmailNotifier) SendEmail(to string, subject string, body string) error {
	// Simulate sending an email
	fmt.Printf("Sending Email to: %s, Subject: %s, Body: %s\n", to, subject, body)
	return nil
}
