package main

import (
	"bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net/http"
)

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

func convertHeader(fasthttpHeader *fasthttp.RequestHeader) http.Header {
	header := make(http.Header)

	fasthttpHeader.VisitAll(func(key, value []byte) {
		header.Set(string(key), string(value))
	})

	return header
}

func main() {

	dbCtx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Create a new Fiber instance
	app := fiber.New()

	// Set up a route to handle incoming requests
	app.All("/*", func(c *fiber.Ctx) error {

		// Check IP address for previous requests
		ip := c.IP()
		currentValue, err := getAndIncrementIPValue(rdb, ip, dbCtx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		body, statusCode, headers, err := makeRequest(c.Method(), "https://www.google.com", c.Body(), convertHeader(&c.Request().Header))
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

	// Start the server on port 7654
	log.Fatal(app.Listen(":7654"))
}
