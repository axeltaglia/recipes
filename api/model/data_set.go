package model

var Ingredients = []Ingredient{
	{ID: 1, Name: "Arroz", Unit: "grs"},
	{ID: 2, Name: "Pechuga de Pollo", Unit: "Kgs"},
	{ID: 3, Name: "Cebolla", Unit: "unid"},
	{ID: 4, Name: "Tomate", Unit: "unid"},
	{ID: 5, Name: "Fideos", Unit: "grs"},
	{ID: 6, Name: "Carne Picada", Unit: "Kgs"},
	{ID: 7, Name: "Queso Rallado", Unit: "grs"},
	{ID: 8, Name: "Leche", Unit: "ml"},
	{ID: 9, Name: "Huevos", Unit: "unid"},
	{ID: 10, Name: "Pan", Unit: "unid"},
}

var Recipes = []Recipe{
	{ID: 1, Name: "Arroz con Pollo", Score: 0},
	{ID: 2, Name: "Fideos con Salsa", Score: 0},
	{ID: 3, Name: "Milanesas con Puré", Score: 0},
	{ID: 4, Name: "Tortilla de Papas", Score: 0},
	{ID: 5, Name: "Hamburguesas Caseras", Score: 0},
	{ID: 6, Name: "Sopa de Verduras", Score: 0},
	{ID: 7, Name: "Pollo al Horno", Score: 0},
	{ID: 8, Name: "Arroz Primavera", Score: 0},
	{ID: 9, Name: "Panqueques", Score: 0},
	{ID: 10, Name: "Ensalada Mixta", Score: 0},
}

var RecipeIngredients = []RecipeIngredient{
	{RecipeID: 1, IngredientID: 1, QuantityPerPerson: 50},
	{RecipeID: 1, IngredientID: 2, QuantityPerPerson: 0.25},
	{RecipeID: 2, IngredientID: 5, QuantityPerPerson: 100},
	{RecipeID: 2, IngredientID: 3, QuantityPerPerson: 1},
	{RecipeID: 2, IngredientID: 4, QuantityPerPerson: 2},
	{RecipeID: 3, IngredientID: 6, QuantityPerPerson: 0.3},
	{RecipeID: 3, IngredientID: 8, QuantityPerPerson: 200},
	{RecipeID: 4, IngredientID: 3, QuantityPerPerson: 1},
	{RecipeID: 4, IngredientID: 9, QuantityPerPerson: 2},
	{RecipeID: 5, IngredientID: 6, QuantityPerPerson: 0.25},
	{RecipeID: 5, IngredientID: 10, QuantityPerPerson: 1},
	{RecipeID: 6, IngredientID: 3, QuantityPerPerson: 1},
	{RecipeID: 6, IngredientID: 4, QuantityPerPerson: 1},
	{RecipeID: 7, IngredientID: 2, QuantityPerPerson: 0.3},
	{RecipeID: 7, IngredientID: 4, QuantityPerPerson: 2},
	{RecipeID: 8, IngredientID: 1, QuantityPerPerson: 60},
	{RecipeID: 8, IngredientID: 4, QuantityPerPerson: 1},
	{RecipeID: 9, IngredientID: 8, QuantityPerPerson: 150},
	{RecipeID: 9, IngredientID: 9, QuantityPerPerson: 1},
	{RecipeID: 10, IngredientID: 4, QuantityPerPerson: 1},
	{RecipeID: 10, IngredientID: 3, QuantityPerPerson: 1},
}

var Steps = []Step{
	{ID: 1, RecipeID: 1, Description: "Lavar el arroz", Order: 1},
	{ID: 2, RecipeID: 1, Description: "Cocinar el pollo", Order: 2},

	{ID: 3, RecipeID: 2, Description: "Hervir los fideos", Order: 1},
	{ID: 4, RecipeID: 2, Description: "Preparar salsa con cebolla y tomate", Order: 2},

	{ID: 5, RecipeID: 3, Description: "Freír milanesas", Order: 1},
	{ID: 6, RecipeID: 3, Description: "Preparar puré de papas", Order: 2},

	{ID: 7, RecipeID: 4, Description: "Batir los huevos", Order: 1},
	{ID: 8, RecipeID: 4, Description: "Freír la mezcla de papa y huevo", Order: 2},

	{ID: 9, RecipeID: 5, Description: "Formar hamburguesas con carne picada", Order: 1},
	{ID: 10, RecipeID: 5, Description: "Cocinar en sartén o plancha", Order: 2},

	{ID: 11, RecipeID: 6, Description: "Picar las verduras", Order: 1},
	{ID: 12, RecipeID: 6, Description: "Hervir con agua y sal", Order: 2},

	{ID: 13, RecipeID: 7, Description: "Condimentar el pollo", Order: 1},
	{ID: 14, RecipeID: 7, Description: "Hornear durante 40 minutos", Order: 2},

	{ID: 15, RecipeID: 8, Description: "Saltear las verduras", Order: 1},
	{ID: 16, RecipeID: 8, Description: "Agregar arroz previamente cocido", Order: 2},

	{ID: 17, RecipeID: 9, Description: "Mezclar huevo y leche", Order: 1},
	{ID: 18, RecipeID: 9, Description: "Cocinar los panqueques en sartén", Order: 2},

	{ID: 19, RecipeID: 10, Description: "Lavar vegetales", Order: 1},
	{ID: 20, RecipeID: 10, Description: "Cortar y servir con aderezo", Order: 2},
}

var RecipeRatings = []RecipeRating{}
