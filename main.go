package main

import (
	"encoding/json"
	"fmt"
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
		//create a map which holds strings as key and any data type as values
		//TODO How does interface work ?
		var jsonStore map[string]interface{}
		//create a pointer to memory location
		var storePtr = &jsonStore
		err := json.Unmarshal(ctx.Body(), storePtr)
		if err != nil {
			panic(err)
		}
		recMapPrinter(jsonStore, 0)
		return nil
	})

	app.Listen(getPort())
}

func recMapPrinter(json map[string]interface{}, indent int) {
	// Create indentation string
	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += "  "
	}
	// Get keys from json
	keys := make([]string, 0, len(json))
	for key := range json {
		keys = append(keys, key)
	}

	// Iterate over sorted keys
	for _, key := range keys {
		value := json[key]
		if data, ok := value.(map[string]interface{}); ok {
			// if value is another map loop over it again
			fmt.Printf("%s%s:\n", indentString, key)
			recMapPrinter(data, indent+1)
		} else if data, ok := value.([]interface{}); ok {
			// if value is an array loop over it
			fmt.Printf("%s%s: \n", indentString, key)
			for _, itemInArray := range data {
				if data, ok := itemInArray.(string); ok {
					// if item in array is a string, print it
					fmt.Printf("%s %v%s\n", indentString, data, " (string)")
				}
				if mp, ok := itemInArray.(map[string]interface{}); ok {
					// if item in array is a map, call function recursively
					recMapPrinter(mp, indent+1)
				}

			}

		} else {
			// if the value is not a map or array, print it
			fmt.Printf("%s%s: %v\n", indentString, key, value)
		}
	}
}
