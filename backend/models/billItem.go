package models

type BillItem struct {
	ID     uint `gorm:"primaryKey"`
	BillID uint	
	Bill Bill `gorm:"foreignKey:BillID"`
	ItemID uint
	Item Item `gorm:"foreignKey:ItemID"`
	Quantity int `gorm:"not null"`
	SubTotal float64 `gorm:"not null"`
}