package models

import (
	"time"
)

type Bill struct {
	ID uint `gorm:"primaryKey"`
	TotalPrice float64 `gorm:"not null"`
	BillItems []BillItem `gorm:"foreignKey:BillID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}