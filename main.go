package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type Wiz struct {
	key      string
	parent   *Wiz
	objType  string
	children *Wiz
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello, I am the JSON Type Wizard!",
		})
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		println(string(ctx.Body()))
		return nil
	})

	app.Listen(getPort())
}
