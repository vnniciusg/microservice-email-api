package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/microservice-email-api/config"
	"github.com/vnniciusg/microservice-email-api/internal/application"
	"github.com/vnniciusg/microservice-email-api/internal/domain"
)

type EmailHandler struct {
	EmailService application.EmailService
}

func NewEmailHandler(emailService application.EmailService) *EmailHandler {
	return &EmailHandler{
		EmailService: emailService,
	}
}

func (e *EmailHandler) SendEmail(c *gin.Context) {
	var email domain.Email

	if err := c.ShouldBindJSON(&email); err != nil {
		restErr := config.NewBadRequestErr("Error")
		c.JSON(restErr.Code, restErr)
		return
	}

	if err := e.EmailService.SendEmail(&email); err != nil {
		restErr := config.NewInternalServerError("Something goes wrong send email")
		c.JSON(restErr.Code, restErr)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully!"})
}
