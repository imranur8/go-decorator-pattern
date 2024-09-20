package notifier

import "fmt"

type SmsNotifier struct {
	BaseNotifier
}

func (s *SmsNotifier) SendSms(to string, message string) error {
	// Simulate sending an SMS
	fmt.Printf("Sending SMS to: %s, Message: %s\n", to, message)
	return nil
}
