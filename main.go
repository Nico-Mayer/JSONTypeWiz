package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
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
		//create a map which holds strings as key and any data type as values
		//TODO How does interface work ?
		var jsonStore map[string]interface{}
		//create a pointer to memory location
		var storePtr = &jsonStore
		err := json.Unmarshal(ctx.Body(), storePtr)
		if err != nil {
			panic(err)
		}
		recMapPrinter(jsonStore)
		return nil
	})

	app.Listen(getPort())
}

func recMapPrinter(json map[string]interface{}) {
	for key, value := range json {
		//if value is another map loop over it again
		if data, ok := value.(map[string]interface{}); ok {
			fmt.Printf("Key: %s, holds nested map \n", key)
			recMapPrinter(data)
		}
		if data, ok := value.([]map[string]interface{}); ok {
			for _, item := range data {
				recMapPrinter(item)
			}
		}
		fmt.Printf("Key: %s, Value: %s \n", key, value)
	}
}
