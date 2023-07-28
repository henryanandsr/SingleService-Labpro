package migrate

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		initializers.LoadEnvVariables("../.env")
	} else {
		os.Getenv("DB_URL")
		os.Getenv("PORT")
	}
}
func MigrateAndSeed() {
	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if !db.Migrator().HasTable(&model.Company{}) {
		err := db.Migrator().CreateTable(&model.Company{})
		if err != nil {
			log.Fatalf("failed to create companies table: %v", err)
		}
		result := db.Create(&model.Company{ID: uuid.New().String(), Nama: "PT. Labpro Indonesia", Alamat: "Jl. Raya ITS", NoTelepon: "031-123456", KodePajak: "ABC"})
		if result.Error != nil {
			fmt.Println("Create statement error:", result.Error)
		} else {
			fmt.Println("Create statement executed successfully")
		}

	}

	if !db.Migrator().HasTable(&model.Barang{}) {
		err := db.Migrator().CreateTable(&model.Barang{})
		if err != nil {
			log.Fatalf("failed to create barang table: %v", err)
		}
		db.Create(&model.Barang{KodeBarang: "B001", NamaBarang: "Buku", HargaBarang: 10000, StokBarang: 10, PerusahaanPembuat: "blablabla"})
	}

	if !db.Migrator().HasTable(&model.User{}) {
		err := db.Migrator().CreateTable(&model.User{})
		if err != nil {
			log.Fatalf("failed to create user table: %v", err)
		}

		db.Create(&model.User{Username: "admin", Password: "admin"})
	}
}
