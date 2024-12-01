package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")

	// Log or use these variables as needed
	log.Println("Database host:", dbHost)
}
