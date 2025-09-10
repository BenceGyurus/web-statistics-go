package database

import (
	"statistics/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func DatabaseInitSession() error {
	dsn := "host=localhost user=root password=12345 dbname=statistics port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Session = db

	err = db.AutoMigrate(&structs.WebMetric{})
	if err != nil {
		return err
	}

	return nil
}
