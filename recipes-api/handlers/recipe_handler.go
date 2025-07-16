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

func findIngredientByID(id int) *model.Ingredient {
	for _, ing := range model.Ingredients {
		if ing.ID == id {
			return &ing
		}
	}
	return nil
}
