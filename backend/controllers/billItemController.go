package controllers

import (
	"backend/initializers"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BillItemGetAll(ctx *gin.Context) {
	var items []models.BillItem
	initializers.DB.Preload("Item").Preload("Bill").Find(&items)

	ctx.JSON(http.StatusOK, gin.H{"bill_items": items})
}

func BillItemCreate(ctx *gin.Context) {
	var body struct {
		BillID uint `json:"bill_id" binding:"required"`
		ItemID uint `json:"item_id" binding:"required"`
		Quantity float64 `json:"quantity" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var bill models.Bill
	if err := initializers.DB.First(&bill, body.BillID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "BillID not found",
		})
	}

	var item models.Item
	if err := initializers.DB.First(&item, body.ItemID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "ItemID not found",
		})
	}

	subTotal := float64(body.Quantity) * item.Price

	billItem := models.BillItem {
		BillID: body.BillID,
		ItemID: body.ItemID,
		Quantity: int(body.Quantity),
		SubTotal: subTotal,
	}

	if err := initializers.DB.Create(&billItem).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	utils.RecalculateBillTotal(body.BillID)

	ctx.JSON(http.StatusOK, gin.H{
		"billItem": billItem,
	})
}

func BillItemUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Quantity int  `json:"quantity"`
		ItemID   uint `json:"item_id"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tìm bill item
	var billItem models.BillItem
	if err := initializers.DB.First(&billItem, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "BillItem not found"})
		return
	}

	// Nếu ItemID thay đổi → kiểm tra tồn tại
	if body.ItemID != 0 {
		var item models.Item
		if err := initializers.DB.First(&item, body.ItemID).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ItemID not found"})
			return
		}

		billItem.ItemID = body.ItemID
		billItem.SubTotal = float64(billItem.Quantity) * item.Price
	}

	// Nếu Quantity thay đổi → update subtotal
	if body.Quantity > 0 {
		billItem.Quantity = body.Quantity

		var item models.Item
		initializers.DB.First(&item, billItem.ItemID)
		billItem.SubTotal = float64(body.Quantity) * item.Price
	}

	initializers.DB.Save(&billItem)

	utils.RecalculateBillTotal(billItem.BillID)

	ctx.JSON(http.StatusOK, gin.H{"bill_item": billItem})
}

func BillItemDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var billItem models.BillItem
	if err := initializers.DB.First(&billItem, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "BillItem not found"})
		return
	}

	if err := initializers.DB.Delete(&models.BillItem{}, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔥 Tự update lại total bill
	utils.RecalculateBillTotal(billItem.BillID)

	ctx.JSON(http.StatusOK, gin.H{"msg": "BillItem deleted"})
}