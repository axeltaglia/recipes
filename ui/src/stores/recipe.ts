import { defineStore } from 'pinia'
import { ref } from 'vue'
// import axios from 'axios'

export interface Recipe {
  id: string
  name: string
  score: number
}

export const useRecipeStore = defineStore('recipe', () => {
  const recipes = ref<Recipe[]>([])

  // agregar async cuando usemos axios
  function fetchRecipes() {
    if (recipes.value.length > 0) return

    recipes.value = [
      { id: '1', name: 'Tarta de manzana', score: 4.5 },
      { id: '2', name: 'Milanesas con puré', score: 5 },
      { id: '3', name: 'Ensalada César', score: 3.8 }
    ]

    // Estos es con axios
    // const { data } = await axios.get<Recipe[]>('/api/v1/recipes')
    // recipes.value = data
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
