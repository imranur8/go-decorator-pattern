package notifier

type Notifier interface {
	SendEmail(to string, subject string, body string) error
	SendSms(to string, message string) error
	SendSlackMessage(channel string, message string) error
}
