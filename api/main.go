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
	api.Get("/recipes/:id", handler.GetRecipeByID)
	api.Post("/recipes/:id/rate", handler.RateRecipe)
	api.Get("/recipes", handler.GetRecipes)

	err := app.Listen(":9123")
	if err != nil {
		log.Fatal(err)
		return
	}
}
