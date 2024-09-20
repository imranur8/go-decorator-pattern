package notifier

type WelcomeEmailNotifier struct {
	EmailNotifier Notifier
	To            string
	Subject       string
	Body          string
}

func NewWelcomeEmailNotifier(welcomeEmailNotifier WelcomeEmailNotifier) *WelcomeEmailNotifier {
	return &WelcomeEmailNotifier{
		EmailNotifier: welcomeEmailNotifier.EmailNotifier,
		To:            welcomeEmailNotifier.To,
		Subject:       welcomeEmailNotifier.Subject,
		Body:          welcomeEmailNotifier.Body,
	}
}

func (g *WelcomeEmailNotifier) SendEmail() error {
	return g.EmailNotifier.SendEmail(g.To, g.Subject, g.Body)
}
