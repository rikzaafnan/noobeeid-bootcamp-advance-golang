package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
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

func LoadConfig(filenames ...string) Config {
	err := godotenv.Overload(filenames...)
	if err != nil {
		log.Println(err)
	}

	cfg := Config{
		App: AppConfig{
			Port: GetEnvString("APP_PORT", ":55555"),
		},
		DB: DBConfig{
			Host: GetEnvString("DB_HOST", "localhost"),
			Port: GetEnvString("DB_HOST", "5432"),
			User: GetEnvString("DB_HOST", "user"),
			Pass: GetEnvString("DB_HOST", ""),
			Name: GetEnvString("DB_HOST", "learn"),
		},
	}
	log.Printf("%+v", cfg)
	return cfg
}

func LoadConfigYaml(filename string) (cfg Config) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Println("error : ", err.Error())
	}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		log.Println("error : ", err.Error())
	}

	return

}

func GetEnvString(key string, fallback string) (val string) {
	val = os.Getenv(key)

	if val == "" {
		val = fallback
	}

	return

}
