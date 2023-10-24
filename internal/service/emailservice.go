package service

import "github.com/vnniciusg/microservice-email-api/internal/domain"

type EmailSender interface {
	SendEmail(email *domain.Email) error
}
