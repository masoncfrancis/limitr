package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupEnvVars() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// If the environment variable is not set, print error message and quit

}

func CheckEnvVar(varName string) bool {
	_, exists := os.LookupEnv(varName)
	if !exists {
		return true
	} else {
		return false
	}
}
