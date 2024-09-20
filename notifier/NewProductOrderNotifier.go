package notifier

type NewProductOrderNotifier struct {
	EmailNotifier Notifier
	SlackNotifier Notifier
	To            string
	Subject       string
	Body          string
	Channel       string
	Message       string
}

func CreateNewProductOrderNotifier(newProductOrderNotifier NewProductOrderNotifier) *NewProductOrderNotifier {
	return &NewProductOrderNotifier{
		EmailNotifier: newProductOrderNotifier.EmailNotifier,
		SlackNotifier: newProductOrderNotifier.SlackNotifier,
		To:            newProductOrderNotifier.To,
		Subject:       newProductOrderNotifier.Subject,
		Body:          newProductOrderNotifier.Body,
		Channel:       newProductOrderNotifier.Channel,
		Message:       newProductOrderNotifier.Message,
	}
}

func (g *NewProductOrderNotifier) SendEmail() error {
	return g.EmailNotifier.SendEmail(g.To, g.Subject, g.Body)
}

func (g *NewProductOrderNotifier) SendSlackMessage() error {
	return g.SlackNotifier.SendSlackMessage(g.Channel, g.Message)
}
