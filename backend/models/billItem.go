package models

type BillItem struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	BillID   uint    `json:"bill_id"`
	Bill     Bill    `json:"bill"`
	ItemID   uint    `json:"item_id"`
	Item     Item    `json:"item"`
	Quantity int     `json:"quantity"`
	SubTotal float64 `json:"sub_total"`
}