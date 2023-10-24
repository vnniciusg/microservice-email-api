package application

import (
	"github.com/vnniciusg/microservice-email-api/internal/domain"
	"github.com/vnniciusg/microservice-email-api/internal/service"
)

type EmailService struct {
	EmailSender service.EmailSender
}

func NewEmailService(emailSender service.EmailSender) *EmailService {
	return &EmailService{
		EmailSender: emailSender,
	}
}

func (e *EmailService) SendEmail(email *domain.Email) error {
	return e.EmailSender.SendEmail(email)
}
