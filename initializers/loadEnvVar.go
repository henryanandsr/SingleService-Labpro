package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(p string) {
	err := godotenv.Load(p)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
