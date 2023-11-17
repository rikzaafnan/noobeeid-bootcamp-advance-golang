package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sesi-8/config"
)

func main() {

	err := config.LoadConfig("./.env")
	if err != nil {
		log.Println("no file .env")
	}

	appPort := config.GetConfigString("APP_PORT", ":4444")

	router := fiber.New()

	files := router.Group("/files")
	{
		files.Post("/upload", Upload)
		files.Post("/download", Upload)
	}

	log.Println("application is running")
	router.Listen(appPort)
}

func Upload(c *fiber.Ctx) error {
	return nil
}

func Download(c *fiber.Ctx) error {
	return nil
}
