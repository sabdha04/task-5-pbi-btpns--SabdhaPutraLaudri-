package models

import (
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primaryKey;not null"`
	Title     string    
	Caption   string    
	PhotoURL  string    
	UserID    uint      `gorm:"index"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
