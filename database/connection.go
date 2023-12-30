package database

import (
	"FinalProject/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=sabdha04 dbname=final port=5432 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	DB.AutoMigrate(&models.User{}, &models.Photo{})
}


