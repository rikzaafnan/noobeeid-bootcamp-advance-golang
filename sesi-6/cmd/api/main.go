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
		log.Println("connection db db gorm not connected....")
	}

	if db != nil {
		log.Println("connection db db gorm connected....")
	}

	dbSQLX, err := database.ConnectSQLXPostgres(config.Cfg.DB)
	if err != nil {
		log.Println("connection db sqlx not connected....")
	}

	if db != nil {
		log.Println("connection db sqlx connected....")
	}

	product.RegisterServiceProduct(router, db, dbSQLX, config.Cfg.DB.ConnectionDBLib)

	router.Listen(config.Cfg.App.Port)

}
