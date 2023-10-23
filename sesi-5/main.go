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

	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	log.Println("server runneingh a port", cfg.App.Port)

}

func loadConfig(filename ...string) (cfg Config, err error) {
	err = godotenv.Load(".app.env", ".db.env")
	if err != nil {
		return
	}

	cfg = Config{
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
	return cfg, nil
}

func getEnvStrinc(key string, fallback string) (vcal string) {
	val := os.Getenv(key)

	if val == "" {
		val = fallback
	}

	return

}
