package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/BeehiveBroadband/limitr/internal/config"
	"github.com/BeehiveBroadband/limitr/internal/database"
	"github.com/BeehiveBroadband/limitr/internal/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const version = "20250325"

// makeRequest makes an HTTP request to the specified URL with the specified method, body, and headers
func makeRequest(method, url string, body []byte, headers http.Header) ([]byte, int, http.Header, error) {
	// Create a new HTTP request to forward the incoming request
	req, err := http.NewRequest(method, url, io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, 0, nil, err
	}

	// Copy headers from the original request
	for key, values := range headers {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Copy response headers
	respHeaders := make(http.Header)
	for key, values := range resp.Header {
		for _, value := range values {
			respHeaders.Set(key, value)
		}
	}

	// Copy response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, nil, err
	}

	return respBody, resp.StatusCode, respHeaders, nil
}

func sendSyslogMsg(msgText string) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		return
	}

	// Example syslog message
	msg := logging.SyslogMessage{
		Priority: 5, // User-level (0-23) + Facility (16-31)
		Version:  version,
		Hostname: hostname,
		AppName:  "limitr",
		Msg:      msgText,
	}

	// Replace with your syslog server address

	hostPort := config.GetSyslogHost() + ":" + config.GetSyslogPort()

	conn, err := logging.ConnectToSyslogServer(hostPort)
	if err != nil {
		log.Fatalf("Error connecting to syslog server: %v", err)
	}
	defer conn.Close()

	err = logging.SendMessage(conn, msg)
	if err != nil {
		log.Fatalf("Error sending syslog message: %v", err)
	}
}

// convertHeader converts a fasthttp header to an http header
func convertHeader(fasthttpHeader *fasthttp.RequestHeader) http.Header {
	header := make(http.Header)

	fasthttpHeader.VisitAll(func(key, value []byte) {
		header.Set(string(key), string(value))
	})

	return header
}

// splitCSV splits a CSV string into an array of strings
func splitCSV(input string) []string {
	return strings.Split(input, ",")
}

// setupAndRunServer sets up the Fiber server and starts it
func setupAndRunServer(rdb *redis.Client, dbCtx context.Context) {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "limitr",
	})

	// Custom middleware to add X-Real-IP to locals
	if config.GetIpHeaderKey() != "" {
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("IP", c.Get(config.GetIpHeaderKey()))
			return c.Next()
		})
	} else {
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("IP", c.IP())
			return c.Next()
		})
	}

	if config.IsEnvVarSet("VERBOSE_MODE") { // Check if verbose mode is enabled
		if config.GetVerboseMode() {
			app.Use(logger.New(logger.Config{
				Format:     "limitr: ${locals:IP} - ${status} - ${method} ${path} ${queryParams} - ${latency}\n",
				TimeFormat: "02-Jan-2006 15:04:05",
				TimeZone:   "America/Denver",
			}))
		}
	}

	if config.IsEnvVarSet("SYSLOG_ENABLED") { // Check if syslog is enabled
		if config.GetSyslogEnabled() {
			app.Use(logger.New(logger.Config{
				Format:     "limitr: ${locals:IP} - ${status} - ${method} ${path} ${queryParams} - ${latency}\n",
				TimeFormat: "02-Jan-2006 15:04:05",
				TimeZone:   "America/Denver",
				Output:     nil,
				Done: func(c *fiber.Ctx, logString []byte) {
					if c.Response().StatusCode() != fiber.StatusOK {
						sendSyslogMsg(string(logString))
					}
				},
			}))
		}

		// Set up a route to handle incoming requests
		app.All("/*", func(c *fiber.Ctx) error {

			ip := c.IP() // Get IP address from request

			// Get IP address from header if env var set
			if config.GetIpHeaderKey() != "" {
				// Assuming the header contains a list of ip addresses, split into an array and get the first one
				listString := c.Get(config.GetIpHeaderKey())
				ip = splitCSV(listString)[0]

				// If the header doesn't exist, use the IP address from the request
				if listString == "" {
					ip = c.IP()
				}
			}

			// Check IP address for previous requests
			shouldRestrict, err := database.CheckIp(rdb, ip, dbCtx, config.GetTimeWindow(), config.GetRateLimit())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}

			// If the IP address has made more requests than allowed, return HTTP code 429
			if shouldRestrict {
				return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests")
			}

			// Query string args are in an array of strings, so we need to join them into a single string of key=value pairs
			queryParams := c.Queries()
			queryString := url.Values{}
			for key, value := range queryParams {
				queryString.Add(key, value)
			}
			encodedQueryString := queryString.Encode()

			url := config.GetForwardUrl() + c.Path()
			if len(queryString) > 0 {
				url += "?" + encodedQueryString
			}
			body, statusCode, headers, err := makeRequest(c.Method(), url, c.Body(), convertHeader(&c.Request().Header))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}

			// Set response headers
			for key, values := range headers {
				for _, value := range values {
					c.Set(key, value)
				}
			}

			// Set response status code
			c.Status(statusCode)

			// Send response body
			return c.Send(body)
		})

		if config.GetUseTls() {
			// Start the server with TLS
			fmt.Printf("Limitr server running on port %s with TLS...\n", config.GetPort())
			err := app.ListenTLS(":"+config.GetPort(), "./ssl/cert.pem", "./ssl/cert.key")
			if err != nil {
				log.Fatalf("Error starting server: %v", err)
			}
			return
		} else {
			// Start the server without TLS
			fmt.Printf("Limitr server running on port %s without TLS...\n", config.GetPort())
			err := app.Listen(":" + config.GetPort())
			if err != nil {
				log.Fatalf("Error starting server: %v", err)
			}
			return
		}
	}
}

// main is the entry point for the application
func main() {

	// Print the version
	fmt.Printf("Limitr (version %s)\n", version)

	// Check for -v flag for version
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		os.Exit(0)
	}

	config.SetupEnvVars() // Load environment variables

	// Create a new Redis client
	dbCtx, rdb := database.CreateDbConn()

	setupAndRunServer(rdb, dbCtx)
}
