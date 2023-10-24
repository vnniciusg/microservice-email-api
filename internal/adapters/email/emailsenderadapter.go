package email

import (
	"github.com/vnniciusg/microservice-email-api/internal/domain"
	"gopkg.in/gomail.v2"
)

type EmailSenderAdapter struct {
	SMTPServer   string
	SMTPPort     int
	SMTPUsername string
	SMTPPassowrd string
}

func NewEmailSenderAdapter(smtpServer string, smtpPort int, smtpUsername string, smtpPassword string) *EmailSenderAdapter {
	return &EmailSenderAdapter{
		SMTPServer:   smtpServer,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassowrd: smtpPassword,
	}
}

func (e *EmailSenderAdapter) SendEmail(email *domain.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.SMTPUsername)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	d := gomail.NewDialer(e.SMTPServer, e.SMTPPort, e.SMTPUsername, e.SMTPPassowrd)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
