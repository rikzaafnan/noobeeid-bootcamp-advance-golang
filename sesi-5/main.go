package main

import (
	"log"
	"os"
)

func main() {

	appPort := os.Getenv("APP_PORT")

	if appPort != "" {
		appPort = ":55555"
	}

	log.Println("server runneingh a port", appPort)

}
