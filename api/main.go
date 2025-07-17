package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	handler "recipes/handlers"
)

func main() {
	fmt.Println("Starting server...")
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/recipes", handler.GetRecipes)
	app.Get("/api/v1/recipes/:id", handler.GetRecipeByID)
	app.Post("/api/v1/recipes/:id/rate", handler.RateRecipe)

	err := app.Listen(":9123")
	if err != nil {
		log.Fatal(err)
		return
	}
}
