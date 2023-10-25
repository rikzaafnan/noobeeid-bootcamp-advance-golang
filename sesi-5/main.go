package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sesi-5/config"
	"sesi-5/database"
)

var envFiles = []string{".app.env", ".db.env"}

func main() {

	//cfg := loadConfig(envFiles...)
	cfg := config.LoadConfigYaml("config.yaml")
	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	log.Println(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// load config
		// diini nanti akan mengembalikan config yang berbeda
		// jika ada perubahan pada file nya
		//cfg := loadConfig(envFiles...)
		cfg := config.LoadConfigYaml("config.yaml")
		json.NewEncoder(w).Encode(cfg)
	})

	log.Printf("%+v", cfg)
	log.Println("server running a port", cfg.App.Port)
	http.ListenAndServe(cfg.App.Port, nil)

}
