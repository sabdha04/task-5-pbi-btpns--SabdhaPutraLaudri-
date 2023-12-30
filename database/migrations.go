package database

import "FinalProject/models"

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.Photo{})
}
