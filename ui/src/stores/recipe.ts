// src/stores/recipe.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

export interface Recipe {
  id: number
  name: string
  score: number
}

export const useRecipeStore = defineStore('recipe', () => {
  const recipes = ref<Recipe[]>([])

  async function fetchRecipes(search?: string) {
    const query = search ? `?search=${encodeURIComponent(search)}` : ''
    const { data } = await axios.get<Recipe[]>(`/api/v1/recipes${query}`)
    recipes.value = data
  }

  function getRecipeById(id: string): Recipe | undefined {
    return recipes.value.find(r => r.id === Number(id))
  }

  return {
    recipes,
    fetchRecipes,
    getRecipeById,
  }
})
