package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"recipes/model"
)

func GetRecipes(c *fiber.Ctx) error {
	search := strings.ToLower(c.Query("search", ""))

	var result []GetRecipesResponse

	for _, recipe := range model.Recipes {
		if search != "" {
			nameMatch := strings.Contains(strings.ToLower(recipe.Name), search)

			ingredientMatch := false
			for _, ri := range model.RecipeIngredients {
				if ri.RecipeID == recipe.ID {
					ingredient := findIngredientByID(ri.IngredientID)
					if ingredient != nil && strings.Contains(strings.ToLower(ingredient.Name), search) {
						ingredientMatch = true
						break
					}
				}
			}

			if !nameMatch && !ingredientMatch {
				continue
			}
		}

		result = append(result, GetRecipesResponse{
			ID:    recipe.ID,
			Name:  recipe.Name,
			Score: recipe.Score,
		})
	}

	return c.JSON(result)
}

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

	return c.JSON(response)
}

func findRecipeByID(recipeID int) *model.Recipe {
	var recipe *model.Recipe
	for _, r := range model.Recipes {
		if r.ID == recipeID {
			recipe = &r
			break
		}
	}
	return recipe
}

func findIngredientByID(id int) *model.Ingredient {
	for _, ing := range model.Ingredients {
		if ing.ID == id {
			return &ing
		}
	}
	return nil
}
