package handler

import (
	"github.com/gofiber/fiber/v2"
	"recipes/model"
)

func GetRecipeByID(c *fiber.Ctx) error {
	recipeID, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	recipe := findRecipeByID(recipeID)

	if recipe == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	var ingredients []RecipeIngredientJSON
	for _, ri := range model.RecipeIngredients {
		if ri.RecipeID != recipeID {
			continue
		}

		for _, ing := range model.Ingredients {
			if ing.ID == ri.IngredientID {
				ingredients = append(ingredients, RecipeIngredientJSON{
					Ingredient: IngredientJSON{
						ID:   ing.ID,
						Name: ing.Name,
						Unit: ing.Unit,
					},
					QuantityPerPerson: ri.QuantityPerPerson,
				})
				break
			}
		}
	}

	var steps []StepJSON
	for _, s := range model.Steps {
		if s.RecipeID == recipeID {
			steps = append(steps, StepJSON{
				ID:          s.ID,
				Description: s.Description,
				Order:       s.Order,
			})
		}
	}

	response := GetRecipeByIDResponse{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Score:       recipe.Score,
		Ingredients: ingredients,
		Steps:       steps,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
