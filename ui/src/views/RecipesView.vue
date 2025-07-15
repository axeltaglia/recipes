<template>
  <section>
    <h1>Listado de Recetas</h1>
    <RecipeList :recipes="recipes" />
  </section>
</template>

<script setup lang="ts">
import RecipeList from '@/components/RecipeList.vue'
import { useRecipeStore } from '@/stores/recipe'
import { storeToRefs } from 'pinia'

const recipeStore = useRecipeStore()
const { recipes } = storeToRefs(recipeStore)

recipeStore.fetchRecipes()
</script>
