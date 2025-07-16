// src/stores/recipe.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

export interface Ingredient {
  id: number
  name: string
  unit: string
}

export interface RecipeIngredient {
  ingredient: Ingredient
  quantityPerPerson: number
}

export interface Step {
  id: number
  description: string
  order: number
}

export interface Recipe {
  id: number
  name: string
  score: number
  ingredients?: RecipeIngredient[]
  steps?: Step[]
}

export const useRecipeStore = defineStore('recipe', () => {
  const recipes = ref<Recipe[]>([])
  const selectedRecipe = ref<Recipe | null>(null)

  async function fetchRecipes(search?: string) {
    const query = search ? `?search=${encodeURIComponent(search)}` : ''
    const { data } = await axios.get<Recipe[]>(`/api/v1/recipes${query}`)
    recipes.value = data
  }

  async function fetchRecipeById(id: number) {
    const { data } = await axios.get<Recipe>(`/api/v1/recipes/${id}`)
    selectedRecipe.value = data
  }

  return {
    recipes,
    selectedRecipe,
    fetchRecipes,
    fetchRecipeById,
  }
})
