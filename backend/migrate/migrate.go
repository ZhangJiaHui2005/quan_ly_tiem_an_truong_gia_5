package main

import (
	"backend/initializers"
	"backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Bill{},
		&models.BillItem{},
		&models.Item{},
		&models.Category{},
	)
}