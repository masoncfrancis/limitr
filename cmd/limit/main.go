package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Set up a route to handle incoming requests
	app.All("/*", func(c *fiber.Ctx) error {
		// Create a new HTTP request to forward the incoming request
		req, err := http.NewRequest(c.Method(), "https://www.google.com", io.NopCloser(bytes.NewReader(c.Body())))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Copy headers from the original request
		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})

		// Perform the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer resp.Body.Close()

		// Copy response headers
		for key, values := range resp.Header {
			for _, value := range values {
				c.Set(key, value)
			}
		}

		// Copy response status code
		c.Status(resp.StatusCode)

		// Copy response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Send(body)
	})

	// Start the server on port 7654
	log.Fatal(app.Listen(":7654"))
}
