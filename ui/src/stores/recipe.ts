import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

export interface Recipe {
  id: string
  name: string
  score: number
}

export const useRecipeStore = defineStore('recipe', () => {
  const recipes = ref<Recipe[]>([])

  async function fetchRecipes() {
    if (recipes.value.length > 0) return

    try {
      const { data } = await axios.get<Recipe[]>('/api/v1/recipes')
      recipes.value = data
    } catch (err) {
      console.error('Error fetching recipes:', err)
    }
  }

  function getRecipeById(id: string): Recipe | undefined {
    return recipes.value.find(r => r.id === id)
  }

  return {
    recipes,
    fetchRecipes,
    getRecipeById,
  }
})
