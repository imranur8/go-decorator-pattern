package notifier

type SignUpEmailNotifier struct {
	EmailNotifier Notifier
	SmsNotifier   Notifier
	To            string
	Subject       string
	Body          string
	Mobile        string
	Message       string
}

func NewSignUpEmailNotifier(signUpEmailNotifier SignUpEmailNotifier) *SignUpEmailNotifier {
	return &SignUpEmailNotifier{
		EmailNotifier: signUpEmailNotifier.EmailNotifier,
		SmsNotifier:   signUpEmailNotifier.SmsNotifier,
		To:            signUpEmailNotifier.To,
		Subject:       signUpEmailNotifier.Subject,
		Body:          signUpEmailNotifier.Body,
		Mobile:        signUpEmailNotifier.Mobile,
		Message:       signUpEmailNotifier.Message,
	}
}

func (g *SignUpEmailNotifier) SendEmail() error {
	return g.EmailNotifier.SendEmail(g.To, g.Subject, g.Body)
}

func (g *SignUpEmailNotifier) SendSms() error {
	return g.SmsNotifier.SendSms(g.Mobile, g.Message)
}
