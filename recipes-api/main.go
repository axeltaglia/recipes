package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Recipe struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

var recipes = []Recipe{
	{ID: 1, Name: "Arroz con pollo", Score: 4.5},
	{ID: 2, Name: "Milanesas con puré", Score: 5},
	{ID: 3, Name: "Ensalada César", Score: 3.8},
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api/v1")

	api.Get("/recipes", func(c *fiber.Ctx) error {
		return c.JSON(recipes)
	})

	app.Listen(":3000")
}
