<template>
  <section>
    <h1>Listado de Recetas</h1>
    <RecipeList :recipes="recipes" />
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import RecipeList from '@/components/RecipeList.vue'
import type { Recipe } from '@/types/Recipe'

const recipes = ref<Recipe[]>([])

onMounted(() => {
  recipes.value = [
    { id: '1', name: 'Tarta de manzana', score: 4.5 },
    { id: '2', name: 'Milanesas con puré', score: 5 },
    { id: '3', name: 'Ensalada César', score: 3.8 }
  ]
})
</script>
