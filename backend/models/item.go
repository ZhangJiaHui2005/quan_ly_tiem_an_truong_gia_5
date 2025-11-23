package models

import (
	"time"
)

type Item struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"not null" json:"name"`
	Price      float64  `gorm:"not null" json:"price"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
