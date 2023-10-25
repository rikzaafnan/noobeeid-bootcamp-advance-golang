package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"net/http"
)

const (
	// config SMTP HOST yaitu gmail
	CONFIG_SMTP_HOST = "smtp.gmail.com"
	// SMTP PORT
	CONFIG_SMTP_PORT = 587
	// Sender Name
	// CONFIG_SENDER_NAME = "Admin NooBeeID <afnangame@gmail.com>"
	CONFIG_SENDER_NAME = "Admin NooBeeID <afnangame@gmail.com>"

	// config untuk authentication nya.
	// gunakan email yg di pakai saat generate app password tadi
	CONFIG_AUTH_EMAIL = "afnangame@gmail.com"
	// app password yang telah di generate
	CONFIG_AUTH_PASSWORD = "APP_PASSWORD"
)

func main() {
	fmt.Println("mail - service")

	router := gin.Default()
	// proses menggunakan middleware
	router.Use(MiddlewareLogging())

	// handler utama
	router.POST("/send-emails", sendEmail)

	router.GET("/ping", func(ctx *gin.Context) {
		log.Println("Masuk ke handler utama")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":4445")
}

func MiddlewareLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("log mail service : incoming request")
		ctx.Next()
		log.Println("log mail service : finish request")
	}
}

func sendEmail(ctx *gin.Context) {
	log.Println("send emails")

	var sendMail SendMail

	err := ctx.ShouldBindJSON(&sendMail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  err.Error(),
		})
		return
	}

	processErr := make(chan error)

	//err = sendMailGoMail(sendMail.To, nil, sendMail.Subject, sendMail.Message, sendMail.From)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "failed send email",
	//		"success": false,
	//		"errors":  err.Error(),
	//	})
	//	return
	//}

	// V2
	go sendMailGoMailV2(sendMail.To, nil, sendMail.Subject, sendMail.Message, sendMail.From, processErr)

	if <-processErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success send email",
		"success": true,
	})
}

type SendMail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

func sendMailGoMail(to []string, cc []string, subject string, message string, from string) (err error) {

	fmt.Println(from)

	// setup gomail message
	mailer := gomail.NewMessage()
	// setting header from
	mailer.SetHeader("From", from)
	// setting header to
	mailer.SetHeader("To", to...)

	// setting header CC
	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}

	// setting subject
	mailer.SetHeader("Subject", subject)
	// setting body
	// kali ini, kita akan menggunakan body HTML agar tampilan dari emailnya lebih menarik
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	return
}

func sendMailGoMailV2(to []string, cc []string, subject string, message string, from string, processErr chan error) {

	// setup gomail message
	mailer := gomail.NewMessage()
	// setting header from
	mailer.SetHeader("From", from)
	// setting header to
	mailer.SetHeader("To", to...)

	// setting header CC
	for _, ccEmail := range cc {
		mailer.SetAddressHeader("Cc", ccEmail, "")
	}

	// setting subject
	mailer.SetHeader("Subject", subject)
	// setting body
	// kali ini, kita akan menggunakan body HTML agar tampilan dari emailnya lebih menarik
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)

	processErr <- err

}
