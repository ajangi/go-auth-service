package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetEnvVariable is a function to get .env variables
func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}