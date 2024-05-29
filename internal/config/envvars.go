package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func SetupEnvVars() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Check if all required environment variables are set
	// Even if a .env file is missing, the environment variables may still be set

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
	if !CheckEnvVar("REDIS_ADDR") {
		log.Println("REDIS_ADDR environment variable is not set, using default (localhost) ")
		err := os.Setenv("REDIS_ADDR", "localhost")
		if err != nil {
			log.Fatal("Error setting default value for REDIS_ADDR")
		}
	}
	if !CheckEnvVar("REDIS_PORT") {
		log.Println("REDIS_PORT environment variable is not set, using default 6379")
		err := os.Setenv("REDIS_PORT", "6379")
		if err != nil {
			log.Fatal("Error setting default value for REDIS_PORT")
		}
	}
	if !CheckEnvVar("REDIS_PASSWORD") {
		log.Println("REDIS_PASSWORD environment variable is not set, using default")
		err := os.Setenv("REDIS_PASSWORD", "")
		if err != nil {
			log.Fatal("Error setting default value for REDIS_PASSWORD")
		}
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

func GetRateLimit() int {
	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		log.Fatal("Error converting RATE_LIMIT to integer")
	}
	return rateLimit
}

func GetTimeWindow() int {
	timeWindow, err := strconv.Atoi(os.Getenv("TIME_WINDOW"))
	if err != nil {
		log.Fatal("Error converting TIME_WINDOW to integer")
	}
	return timeWindow
}
