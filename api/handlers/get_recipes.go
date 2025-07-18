package handler

import (
	"github.com/gofiber/fiber/v2"
	"recipes/model"
	"sort"
	"strings"
)

func GetRecipes(c *fiber.Ctx) error {
	search := strings.ToLower(c.Query("search", ""))
	limit := c.QueryInt("limit", 0)
	orderByScore := c.Query("order_by_score", "")

	if orderByScore == "asc" {
		sort.Slice(model.Recipes, func(i, j int) bool {
			return model.Recipes[i].Score < model.Recipes[j].Score
		})
	} else if orderByScore == "desc" {
		sort.Slice(model.Recipes, func(i, j int) bool {
			return model.Recipes[i].Score > model.Recipes[j].Score
		})
	}

	var result []GetRecipesResponse
	var count int

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

		count++
		if limit > 0 && count == limit {
			break
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
