package main

import (
	"backend/controllers"
	"backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
  initializers.ConnectDatabase()
}

func main() {
  router := gin.Default()

  router.GET("/api/categories", controllers.CategoryGetAll)
  router.POST("/api/categories/create", controllers.CategoryCreate)
  router.PUT("/api/categories/update/:id", controllers.CategoryUpdate)
  router.DELETE("/api/categories/delete/:id", controllers.CategoryDelete)

  router.GET("/api/items", controllers.ItemGetAll)
  router.POST("/api/items/create", controllers.ItemsCreate)
  router.PUT("/api/items/update/:id", controllers.ItemUpdate)
  router.DELETE("/api/items/delete/:id", controllers.ItemDelete)

  router.GET("/api/billitems", controllers.BillItemGetAll)
  router.POST("/api/billitems/create", controllers.BillItemCreate)
  router.PUT("/api/billitems/update/:id", controllers.BillItemUpdate)
  router.DELETE("/api/billitems/delete/:id", controllers.BillItemDelete)

  router.GET("/api/bills", controllers.BillGetAll)
  router.POST("/api/bills/create", controllers.BillCreate)

  router.Run()
}