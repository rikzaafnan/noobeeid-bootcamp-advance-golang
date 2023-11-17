package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"sesi-8/config"
)

var path string

func main() {

	err := config.LoadConfig(".env")
	if err != nil {
		log.Println("no file .env")
	}

	appPort := config.GetConfigString("APP_PORT", ":4444")
	path = config.GetConfigString("PATH", "public/uploads")

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

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("error when try to parse FormFile with detail : ", err.Error())
	}

	if file.Size > 1*1024*1024 {
		errBadRequest := fiber.ErrBadRequest
		errBadRequest.Message = "file to big, expected 1 Mb"
		return errBadRequest
	}

	os.Create(path + "/" + file.Filename)

	return nil
}

func Download(c *fiber.Ctx) error {
	return nil
}
