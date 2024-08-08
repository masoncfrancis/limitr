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
		fmt.Println("No .env file found, checking environment variables...")
	}

	// Check if all required environment variables are set
	// Even if a .env file is missing, the environment variables may still be set

	errorHappened := false

	// If the environment variable is not set, print error message and quit
	if !IsEnvVarSet("PORT") {
		fmt.Println("PORT is not set, using default (7654)")
		err := os.Setenv("PORT", "7654")
		if err != nil {
			fmt.Println("Error setting default value for PORT")
		}
	} else {
		// Validate port
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Println("Error converting PORT to integer")
			errorHappened = true
		} else if port < 1 || port > 65535 {
			fmt.Println("PORT must be between 1 and 65535")
			errorHappened = true
		} else {
			fmt.Println("PORT is set to " + os.Getenv("PORT"))
		}
	}
	if !IsEnvVarSet("FORWARD_URL") {
		fmt.Println("FORWARD_URL is not set and is required")
		errorHappened = true
	} else {
		fmt.Println("FORWARD_URL is set to " + os.Getenv("FORWARD_URL"))
	}
	if !IsEnvVarSet("RATE_LIMIT") {
		fmt.Println("RATE_LIMIT is not set and is required")
		errorHappened = true
	} else {
		rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
		if err != nil {
			fmt.Println("Error converting RATE_LIMIT to integer")
			errorHappened = true
		} else if rateLimit <= 0 {
			fmt.Println("RATE_LIMIT must be greater than 0")
			errorHappened = true
		} else {
			fmt.Println("RATE_LIMIT is set to " + os.Getenv("RATE_LIMIT"))
		}
	}
	if !IsEnvVarSet("TIME_WINDOW") {
		fmt.Println("TIME_WINDOW is not set and is required")
		errorHappened = true
	} else {
		timeWindow, err := strconv.Atoi(os.Getenv("TIME_WINDOW"))
		if err != nil {
			fmt.Println("Error converting TIME_WINDOW to integer")
			errorHappened = true
		} else if timeWindow <= 0 {
			fmt.Println("TIME_WINDOW must be greater than 0")
			errorHappened = true
		} else {
			fmt.Println("TIME_WINDOW is set to " + os.Getenv("TIME_WINDOW"))
		}

	}
	if !IsEnvVarSet("REDIS_ADDR") {
		fmt.Println("REDIS_ADDR is not set, using default (localhost:6379) ")
		err := os.Setenv("REDIS_ADDR", "localhost:6379")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_ADDR")
		}
	}
	if !IsEnvVarSet("REDIS_PASSWORD") {
		fmt.Println("REDIS_PASSWORD is not set, using default (blank)")
		err := os.Setenv("REDIS_PASSWORD", "")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_PASSWORD")
		}
	}

	if !IsEnvVarSet("REDIS_DB") {
		fmt.Println("REDIS_DB is not set, using default (0)")
		err := os.Setenv("REDIS_DB", "0")
		if err != nil {
			fmt.Println("Error setting default value for REDIS_DB")
		}
	} else {
		redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			fmt.Println("Error converting REDIS_DB to integer")
			errorHappened = true
		} else if redisDb < 0 {
			fmt.Println("REDIS_DB must be greater than or equal to 0")
			errorHappened = true
		} else {
			fmt.Println("REDIS_DB is set to " + os.Getenv("REDIS_DB"))
		}
	}

	if !IsEnvVarSet("USE_TLS") {
		fmt.Println("USE_TLS is not set, using default (false)")
		err := os.Setenv("USE_TLS", "false")
		if err != nil {
			fmt.Println("Error setting default value for USE_TLS")
		}
	}

	if !IsEnvVarSet("IP_HEADER_KEY") {
		fmt.Println("IP_HEADER_KEY is not set, using default (blank)")
		err := os.Setenv("IP_HEADER_KEY", "")
		if err != nil {
			fmt.Println("Error setting default value for IP_HEADER_KEY")
		}
	} else {
		fmt.Println("IP_HEADER_KEY is set to " + os.Getenv("IP_HEADER_KEY"))
	}

	if !IsEnvVarSet("VERBOSE_MODE") {
		fmt.Println("VERBOSE_MODE is not set, using default (false)")
		err := os.Setenv("VERBOSE_MODE", "false")
		if err != nil {
			fmt.Println("Error setting default value for VERBOSE_MODE")
		}
	}
	if IsEnvVarSet("VERBOSE_MODE") {
		fmt.Println("VERBOSE_MODE is set to " + os.Getenv("VERBOSE_MODE"))
	}

	// Check if SYSLOG_ENABLED is set, if so check syslog host and port
	if !IsEnvVarSet("SYSLOG_ENABLED") {
		fmt.Println("SYSLOG_ENABLED is not set, using default (false)")
		err := os.Setenv("SYSLOG_ENABLED", "false")
		if err != nil {
			fmt.Println("Error setting default value for SYSLOG_ENABLED")
		}
	} else {
		fmt.Println("SYSLOG_ENABLED is set to " + os.Getenv("SYSLOG_ENABLED"))
		if os.Getenv("SYSLOG_ENABLED") == "true" {
			if !IsEnvVarSet("SYSLOG_HOST") {
				fmt.Println("SYSLOG_HOST is not set, is required when SYSLOG_ENABLED is true")
				errorHappened = true
			} else {
				fmt.Println("SYSLOG_HOST is set to " + os.Getenv("SYSLOG_HOST"))
			}
			if !IsEnvVarSet("SYSLOG_PORT") {
				fmt.Println("SYSLOG_PORT is not set, is required when SYSLOG_ENABLED is true")
				errorHappened = true
			} else {
				fmt.Println("SYSLOG_PORT is set to " + os.Getenv("SYSLOG_PORT"))
			}
		}
	}

	// Exit if any errors occurred
	if errorHappened {
		fmt.Println("Exiting due to missing or improperly set environment variables...")
		os.Exit(1)
	}

}

func IsEnvVarSet(varName string) bool {
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

func GetRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

func GetUseTls() bool {
	useTls, err := strconv.ParseBool(os.Getenv("USE_TLS"))
	if err != nil {
		fmt.Println("Error converting USE_TLS to boolean")
	}
	return useTls

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

func GetRedisDb() int {
	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		fmt.Println("Error converting REDIS_DB to integer")
	}
	return redisDb
}

func GetIpHeaderKey() string {
	return os.Getenv("IP_HEADER_KEY")
}

func GetVerboseMode() bool {
	verboseMode, err := strconv.ParseBool(os.Getenv("VERBOSE_MODE"))
	if err != nil {
		fmt.Println("Error converting VERBOSE_MODE to boolean")
	}
	return verboseMode
}

func GetSyslogEnabled() bool {
	syslogEnabled, err := strconv.ParseBool(os.Getenv("SYSLOG_ENABLED"))
	if err != nil {
		fmt.Println("Error converting SYSLOG_ENABLED to boolean")
	}
	return syslogEnabled
}

func GetSyslogHost() string {
	return os.Getenv("SYSLOG_HOST")
}

func GetSyslogPort() string {
	return os.Getenv("SYSLOG_PORT")
}
