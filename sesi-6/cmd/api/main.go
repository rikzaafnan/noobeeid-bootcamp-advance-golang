package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sesi-6/app/product"
	"sesi-6/config"
	"sesi-6/pkg/database"
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

	db, err := database.ConnectGORMPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected....")
	}

	product.RegisterServiceProduct(router, db)

	router.Listen(config.Cfg.App.Port)

}
