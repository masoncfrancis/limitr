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

	errorHappened := false

	// If the environment variable is not set, print error message and quit
	if !CheckEnvVar("FORWARD_URL") {
		log.Fatal("FORWARD_URL environment variable is not set")
		errorHappened = true
	}
	if !CheckEnvVar("RATE_LIMIT") {
		log.Fatal("RATE_LIMIT environment variable is not set")
		errorHappened = true
	}
	if !CheckEnvVar("TIME_WINDOW") {
		log.Fatal("TIME_WINDOW environment variable is not set")
		errorHappened = true
	}

	if errorHappened {
		log.Println("Exiting due to missing environment variables...")
		os.Exit(1)
	}

}

func CheckEnvVar(varName string) bool {
	_, exists := os.LookupEnv(varName)
	if exists {
		return true
	} else {
		return false
	}
}
