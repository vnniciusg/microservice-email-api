package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/microservice-email-api/config"
	"github.com/vnniciusg/microservice-email-api/internal/adapters/email"
	"github.com/vnniciusg/microservice-email-api/internal/adapters/http/rest"
	"github.com/vnniciusg/microservice-email-api/internal/application"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	emailSenderAdapter := email.NewEmailSenderAdapter(config.SMTPServer, config.SMTPPort, config.SMTPUsername, config.SMTPPassowrd)

	emailService := application.NewEmailService(emailSenderAdapter)

	emailHandler := rest.NewEmailHandler(*emailService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/send-email", emailHandler.SendEmail)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}

}
