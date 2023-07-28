package main

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"log"
	"os"
)

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		initializers.LoadEnvVariables("../.env")
	} else {
		os.Getenv("DB_URL")
		os.Getenv("PORT")
	}
}

func main() {
	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.Company{}, &model.Barang{}, &model.User{})
}
