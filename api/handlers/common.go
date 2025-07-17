package handler

import "recipes/model"

func findRecipeByID(recipeID int) *model.Recipe {
	for _, r := range model.Recipes {
		if r.ID == recipeID {
			return &r
		}
	}
	return nil
}

func findIngredientByID(id int) *model.Ingredient {
	for _, ing := range model.Ingredients {
		if ing.ID == id {
			return &ing
		}
	}
	return nil
}
