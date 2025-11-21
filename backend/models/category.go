package models

import (
	"time"
)

type Category struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"not null;unique" json:"name"`
    Items     []Item    `gorm:"foreignKey:CategoryID" json:"items,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
