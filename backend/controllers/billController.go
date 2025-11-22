package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BillGetAll(ctx *gin.Context) {
	var bills models.Bill

	initializers.DB.Preload("BillItem").Find(&bills)

	ctx.JSON(http.StatusOK, gin.H{
		"bills": bills,
	})
}

func BillCreate(ctx *gin.Context) {
	bill := models.Bill{
		TotalPrice: 0,
	}

	if err := initializers.DB.Create(&bill).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"bill": bill})
}