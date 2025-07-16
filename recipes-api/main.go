package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	handler "recipes/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/recipes", handler.GetRecipes)
	app.Get("/api/v1/recipes/:id", handler.GetRecipeByID)

	app.Listen(":3000")
}
