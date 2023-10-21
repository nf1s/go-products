package main

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	GetProduct(c *fiber.Ctx) error
}

type productController struct {
}

func NewProductController() ProductController {
	return &productController{}
}

func (p *productController) GetProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
