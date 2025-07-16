<template>
  <section>
    <h1>Listado de Recetas</h1>

    <form @submit.prevent="handleSearch" class="search-form">
      <input v-model="searchTerm" type="text" placeholder="Receta..." />
      <button type="submit">Buscar</button>
    </form>

    <RecipeList :recipes="recipes" />
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import RecipeList from '@/components/RecipeList.vue'
import { useRecipeStore } from '@/stores/recipe'
import { storeToRefs } from 'pinia'

const recipeStore = useRecipeStore()
const { recipes } = storeToRefs(recipeStore)

const searchTerm = ref('')

recipeStore.fetchRecipes()

function handleSearch() {
  recipeStore.fetchRecipes(searchTerm.value)
}
</script>

<style scoped>
.search-form {
  margin-bottom: 1rem;
  display: flex;
  gap: 0.5rem;
}

input[type="text"] {
  padding: 0.5rem;
  border: 1px solid #ccc;
  font-size: 1rem;
  flex: 1;
}

button {
  padding: 0.5rem 1rem;
  background: #eee;
  border: 1px solid #ccc;
  font-size: 1rem;
  cursor: pointer;
}
</style>
