package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// setup struct coinfig
type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func main() {

	err := godotenv.Load(".app.env", ".db.env")
	if err != nil {
		panic(err)
	}

	cfg := Config{
		App: AppConfig{
			Port: os.Getenv("APP_PORT"),
		},
		DB: DBConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}

	log.Printf("%+v", cfg)

	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		appPort = ":55555"
	}

	log.Println("server runneingh a port", appPort)

}
