package main

import (
	"notifier-go/notifier"
)

func main() {
	emailNotifier := &notifier.EmailNotifier{}
	smsNotifier := &notifier.SmsNotifier{}
	slackNotifier := &notifier.SlackNotifier{}

	welcomeEmailNotifier := notifier.NewWelcomeEmailNotifier(notifier.WelcomeEmailNotifier{
		EmailNotifier: emailNotifier,
		To:            "example@example.com",
		Body:          "Welcome to our platform!",
		Subject:       "Welcome",
	})

	welcomeEmailNotifier.SendEmail()

	signUpNotifier := notifier.NewSignUpEmailNotifier(notifier.SignUpEmailNotifier{
		EmailNotifier: emailNotifier,
		SmsNotifier:   smsNotifier,
		To:            "example@example.com",
		Body:          "SignUp success!!",
		Subject:       "SignUp",
		Mobile:        "+1234567890",
		Message:       "signup success to our service!",
	})

	signUpNotifier.SendEmail()
	signUpNotifier.SendSms()

	newProductOrderNotifier := notifier.CreateNewProductOrderNotifier(notifier.NewProductOrderNotifier{
		EmailNotifier: emailNotifier,
		SlackNotifier: slackNotifier,
		To:            "example@example.com",
		Body:          "Here is the new Product Order Details",
		Subject:       "Order Submitted",
		Channel:       "#product-order",
		Message:       "Awesome new product order submitted!",
	})

	newProductOrderNotifier.SendEmail()
	newProductOrderNotifier.SendSlackMessage()
}
