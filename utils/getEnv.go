package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVariable - get environment variable by key
func GetEnvVariable(KEY string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error when loading .env file")
	}
	return os.Getenv(KEY)
}
