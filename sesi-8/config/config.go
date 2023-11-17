package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig(filename string) (err error) {

	err = godotenv.Load(filename)
	return
}
func GetConfigString(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}
