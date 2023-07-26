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
var dbError error // Add this line

func GetDBInstance() (*gorm.DB, error) { // Modify this line
	once.Do(func() {
		dsn := os.Getenv("DB_URL")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("failed to connect to database:", err) // Modify this line
			dbError = err                                      // Add this line
			return
		}

		instance = &database{DB: db}
		log.Println("database connection successful")
	})

	if instance == nil {
		return nil, dbError // Add this line
	}
	return instance.DB, nil // Modify this line
}
