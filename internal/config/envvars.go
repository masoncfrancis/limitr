package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func SetupEnvVars() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	// Check if all required environment variables are set
	// Even if a .env file is missing, the environment variables may still be set

	errorHappened := false

	// If the environment variable is not set, print error message and quit
	if !CheckEnvVar("PORT") {
		fmt.Println("PORT environment variable is not set, using default (7654)")
		err := os.Setenv("PORT", "7654")
		if err != nil {
			fmt.Println("Error setting default value for PORT")
		}
	}
	if !CheckEnvVar("FORWARD_URL") {
		fmt.Println("FORWARD_URL environment variable is not set")
		errorHappened = true
	}
	if !CheckEnvVar("RATE_LIMIT") {
		fmt.Println("RATE_LIMIT environment variable is not set")
		errorHappened = true
	}
	if !CheckEnvVar("TIME_WINDOW") {
		fmt.Println("TIME_WINDOW environment variable is not set")
		errorHappened = true
	}
	if !CheckEnvVar("REDIS_ADDR") {
		fmt.Println("REDIS_ADDR environment variable is not set, using default (localhost) ")
		err := os.Setenv("REDIS_ADDR", "localhost")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_ADDR")
		}
	}
	if !CheckEnvVar("REDIS_PORT") {
		fmt.Println("REDIS_PORT environment variable is not set, using default (6379)")
		err := os.Setenv("REDIS_PORT", "6379")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_PORT")
		}
	}
	if !CheckEnvVar("REDIS_PASSWORD") {
		fmt.Println("REDIS_PASSWORD environment variable is not set, using default (blank)")
		err := os.Setenv("REDIS_PASSWORD", "")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_PASSWORD")
		}
	}

	if errorHappened {
		fmt.Println("Exiting due to missing environment variables...")
		os.Exit(1)
	}

}

func CheckEnvVar(varName string) bool {
	_, exists := os.LookupEnv(varName)
	if exists {
		return true // Environment variable is set
	} else {
		return false // Environment variable is not set
	}
}

// Getters

func GetForwardUrl() string {
	return os.Getenv("FORWARD_URL")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}

func GetRedisPort() string {
	return os.Getenv("REDIS_PORT")
}

func GetRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

func GetRateLimit() int {
	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		fmt.Println("Error converting RATE_LIMIT to integer")
	}
	return rateLimit
}

func GetTimeWindow() int {
	timeWindow, err := strconv.Atoi(os.Getenv("TIME_WINDOW"))
	if err != nil {
		fmt.Println("Error converting TIME_WINDOW to integer")
	}
	return timeWindow
}
