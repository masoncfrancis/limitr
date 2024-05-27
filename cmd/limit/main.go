package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This was served with Go Fiber")
	})

	fmt.Println("Starting server at port 7654...")

	app.Listen(":7654")
}
