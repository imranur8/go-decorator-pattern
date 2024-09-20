package notifier

import "errors"

type BaseNotifier struct{}

func (b *BaseNotifier) SendEmail(to string, subject string, body string) error {
	return errors.New("SendEmail not implemented")
}

func (b *BaseNotifier) SendSms(to string, message string) error {
	return errors.New("SendSms not implemented")
}

func (b *BaseNotifier) SendSlackMessage(channel string, message string) error {
	return errors.New("SendSlackMessage not implemented")
}
