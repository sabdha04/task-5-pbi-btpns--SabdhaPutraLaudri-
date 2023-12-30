package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;not null"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
