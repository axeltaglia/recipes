package model

type Ingredient struct {
	ID   int
	Name string
	Unit string
}

type RecipeIngredient struct {
	IngredientID      int
	RecipeID          int
	QuantityPerPerson float64
}

type Step struct {
	ID          int
	RecipeID    int
	Description string
	Order       int
}

type Recipe struct {
	ID    int
	Name  string
	Score float64
}

type RecipeRating struct {
	RecipeID int
	Value    int
}
