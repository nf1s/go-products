package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	controller := NewProductController()

	app := fiber.New()
	app.Get("/", controller.GetProduct)
	app.Listen(":8000")
}
