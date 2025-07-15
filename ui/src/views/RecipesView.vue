<template>
  <section>
    <h1>Listado de Recetas</h1>
    <RecipeList :recipes="recipes" />
  </section>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import RecipeList from '@/components/RecipeList.vue'
import { useRecipeStore } from '@/stores/recipe'

const recipeStore = useRecipeStore()

onMounted(() => {
  recipeStore.fetchRecipes()
})

const recipes = recipeStore.recipes
</script>
