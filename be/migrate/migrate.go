package main

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables("../.env")
}

func main() {
	err := initializers.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	initializers.DB.AutoMigrate(&model.Company{}, &model.Barang{}, &model.User{})
}
