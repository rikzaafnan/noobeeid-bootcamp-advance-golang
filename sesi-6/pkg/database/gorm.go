package database

import (
	"fmt"
	"gorm.io/gorm"
	"sesi-6/config"
)

func ConnectGORMPostgres(cfg config.DB) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)

	gorm.Open()

	return
}
