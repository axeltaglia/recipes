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

  async function fetchRecipes(search?: string, orderByScore?: string, limit?: number) {
    const params = new URLSearchParams()

    if (search) {
      params.append('search', search)
    }
    if (orderByScore) {
      params.append('order_by_score', orderByScore)
    }
    if (limit && limit > 0) {
      params.append('limit', limit.toString())
    }

    const queryString = params.toString()
    const { data } = await axios.get<Recipe[]>(`/api/v1/recipes${queryString ? '?' + queryString : ''}`)
    recipes.value = data
  }

  async function fetchRecipeById(id: number) {
    const { data } = await axios.get<Recipe>(`/api/v1/recipes/${id}`)
    selectedRecipe.value = data
  }

  async function rateRecipe(id: number, value: number) {
    const { data } = await axios.post<{ score: number }>(`/api/v1/recipes/${id}/rate`, { value })

    const recipe = recipes.value.find(r => r.id === id)
    if (recipe) {
      recipe.score = data.score
    }

    if (selectedRecipe.value && selectedRecipe.value.id === id) {
      selectedRecipe.value.score = data.score
    }
  }

  return {
    recipes,
    selectedRecipe,
    fetchRecipes,
    fetchRecipeById,
    rateRecipe
  }
})
