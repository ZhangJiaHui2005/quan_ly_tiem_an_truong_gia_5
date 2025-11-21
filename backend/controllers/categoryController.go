package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryGetAll(ctx *gin.Context) {
	var categories []models.Category

	initializers.DB.Find(&categories)

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func CategoryCreate(ctx *gin.Context) {
	var body struct {
		Name string
	}

	ctx.ShouldBindJSON(&body)

	category := models.Category{
		Name: body.Name,
	}

	result := initializers.DB.Create(&category)

	if result.Error != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func CategoryUpdate(ctx *gin.Context) {
    id := ctx.Param("id")

    var body struct {
        Name string `json:"name"`
    }

    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Tìm category
    var category models.Category
    if err := initializers.DB.First(&category, id).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "msg": "No category found",
        })
        return
    }

    // Cập nhật
    if err := initializers.DB.Model(&category).Updates(models.Category{
        Name: body.Name,
    }).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "category": category,
    })
}

func CategoryDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	var category models.Category
	if err := initializers.DB.First(&category, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "Category not found",
		})
		return
	}

	if err := initializers.DB.Delete(&category, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Category deleted successfully",
	})		
}