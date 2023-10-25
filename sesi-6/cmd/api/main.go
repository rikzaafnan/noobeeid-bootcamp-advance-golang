package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sesi-6/config"
)

func main() {

	router := fiber.New(fiber.Config{
		AppName: "Product Services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error : ", err.Error())
	}

	router.Listen(config.Cfg.App.Port)

}
