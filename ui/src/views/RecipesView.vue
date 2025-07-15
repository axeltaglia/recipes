<template>
  <section>
    <h1>Listado de Recetas</h1>
    <RecipeList :recipes="recipes" />
  </section>
</template>

<script setup lang="ts">
import RecipeList from '@/components/RecipeList.vue'

interface Recipe {
  name: string
  score: number
}

const recipes: Recipe[] = [
  { name: 'Tarta de manzana', score: 4.5 },
  { name: 'Milanesas con puré', score: 5 },
  { name: 'Ensalada César', score: 3.8 }
]
</script>
