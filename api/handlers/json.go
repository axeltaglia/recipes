package handler

type GetRecipesResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type GetRecipeByIDResponse struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Score       float64                `json:"score"`
	Ingredients []RecipeIngredientJSON `json:"ingredients"`
	Steps       []StepJSON             `json:"steps"`
}

type RecipeIngredientJSON struct {
	Ingredient        IngredientJSON `json:"ingredient"`
	QuantityPerPerson float64        `json:"quantityPerPerson"`
}

type IngredientJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

type StepJSON struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

type RateRecipeRequest struct {
	Value int `json:"value"`
}

type RateRecipeResponse struct {
	RecipeID int     `json:"recipe_id"`
	Score    float64 `json:"score"`
}
