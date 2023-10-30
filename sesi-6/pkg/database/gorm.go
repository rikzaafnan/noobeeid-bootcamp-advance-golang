package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sesi-6/config"
)

func ConnectGORMPostgres(cfg config.DB) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	//db.Debug().AutoMigrate(product.Product{})
	db.Debug()

	return
}
