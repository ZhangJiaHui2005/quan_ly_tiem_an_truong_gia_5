package utils

import (
	"backend/initializers"
	"backend/models"
)

func RecalculateBillTotal(billID uint) error {
	var total float64

	initializers.DB.Model(&models.BillItem{}).Where("bill_id = ?", billID).Select("Sum(sub_total)").Scan(&total)

	return initializers.DB.Model(&models.Bill{}).
		Where("id = ?", billID).
		Update("total_price", total).Error
}