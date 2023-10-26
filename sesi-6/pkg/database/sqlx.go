package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sesi-6/config"
)

func ConnectSQLXPostgres(cfg config.DB) (db *sqlx.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	return
}
