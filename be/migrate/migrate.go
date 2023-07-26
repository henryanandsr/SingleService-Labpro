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
	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.Company{}, &model.Barang{}, &model.User{})
}
