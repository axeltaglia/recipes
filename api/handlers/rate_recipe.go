package handler

import (
	"github.com/gofiber/fiber/v2"
	"recipes/model"
)

func RateRecipe(c *fiber.Ctx) error {
	recipeID, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var recipe *model.Recipe
	for i := range model.Recipes {
		if model.Recipes[i].ID == recipeID {
			recipe = &model.Recipes[i]
			break
		}
	}
	if recipe == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	var req RateRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}
	if req.Value < 1 || req.Value > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Rating must be between 1 and 5"})
	}

	model.RecipeRatings = append(model.RecipeRatings, model.RecipeRating{
		RecipeID: recipeID,
		Value:    req.Value,
	})

	var total, count int
	for _, r := range model.RecipeRatings {
		if r.RecipeID == recipeID {
			total += r.Value
			count++
		}
	}
	if count > 0 {
		recipe.Score = float64(total) / float64(count)
	}

	res := RateRecipeResponse{
		RecipeID: recipeID,
		Score:    recipe.Score,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
