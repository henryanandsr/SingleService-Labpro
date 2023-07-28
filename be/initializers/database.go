package initializers

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	DB *gorm.DB
}

var instance *database
var once sync.Once
var dbError error

func GetDBInstance() (*gorm.DB, error) {
	once.Do(func() {
		dsn := os.Getenv("DB_URL")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("failed to connect to database:", err)
			dbError = err                                     
			return
		}

		instance = &database{DB: db}
		log.Println("database connection successful")
	})

	if instance == nil {
		return nil, dbError
	}
	return instance.DB, nil
}
