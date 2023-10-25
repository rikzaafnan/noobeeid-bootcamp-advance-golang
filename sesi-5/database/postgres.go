package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sesi-5/config"
)

func ConnectPostgres(dbConfig config.DBConfig) (db *sql.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Pass, dbConfig.Name)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {

		return

	}

	return
}
