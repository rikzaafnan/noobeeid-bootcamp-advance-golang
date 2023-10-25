package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	httpClient "mail-smtp/http-client"
	"net/http"
)

const SEND_EMAIL_SERVICE_URL = "http://localhost:4445"

func main() {
	log.Println("app-service")

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://bootcamp.noobee.id"} // Ganti dengan domain Anda
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// proses menggunakan middleware
	router.Use(cors.New(config))
	router.Use(MiddlewareLogging())

	// handler utama
	router.POST("/send", sendEmail)

	router.GET("/ping", func(ctx *gin.Context) {
		log.Println("Masuk ke handler utama")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":4444")

}

func MiddlewareLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("log App service : incoming request")
		ctx.Next()
		log.Println("log App service : finish request")
	}
}

func sendEmail(ctx *gin.Context) {
	log.Println("Send Email")

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

	// init client
	client := httpClient.NewHttpClient(SEND_EMAIL_SERVICE_URL)

	// proses mengirim request dengan method POST
	resp, err := client.Post("/send-emails", sendMail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  err.Error(),
		})
		return
	}

	responseString := string(resp)

	if responseString == "404 page not found" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  responseString,
		})
		return
	}

	var responseSendEmail ResponseSendEmail

	err = json.Unmarshal(resp, &responseSendEmail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  err.Error(),
		})
		return
	}

	if responseSendEmail.Success == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed send email",
			"success": false,
			"errors":  responseSendEmail.Message,
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

type ResponseSendEmail struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
