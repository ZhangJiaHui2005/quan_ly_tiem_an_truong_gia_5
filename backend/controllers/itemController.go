package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ItemGetAll(ctx *gin.Context) {
	var items []models.Item

	initializers.DB.Find(&items)

	ctx.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func ItemsCreate(ctx *gin.Context) {
	var body struct {
		Name       string  `json:"name" binding:"required"`
		Price      float64 `json:"price" binding:"required"`
		CategoryID uint    `json:"category_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var category models.Category
    if err := initializers.DB.First(&category, body.CategoryID).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "CategoryID not found",
        })
        return
    }

	item := models.Item{
		Name: body.Name,
		Price: body.Price,
		CategoryID: body.CategoryID,
	}

	if err := initializers.DB.Create(&item).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func ItemUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Name       string  `json:"name" binding:"required"`
		Price      float64 `json:"price" binding:"required"`
		CategoryID uint    `json:"category_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var item models.Item

	if err := initializers.DB.First(&item, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err := initializers.DB.Model(&item).Updates(&models.Item{
		Name: body.Name,
		Price: body.Price,
		CategoryID: body.CategoryID,
	})

	if err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func ItemDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := initializers.DB.Delete(&models.Item{}, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Item deleted"})
}