<template>
  <section>
    <h1>Valoraciones</h1>

    <div class="tabs">
      <button
        :class="{ active: activeTab === 'bestRated'}"
        @click="activeTab = 'bestRated'"
      >
        Mejor Puntaje
      </button>
      <button
        :class="{ active: activeTab === 'worstRated'}"
        @click="activeTab = 'worstRated'"
      >
        Peor Puntaje
      </button>
    </div>

    <div class="tab-content">
      <div v-if="activeTab === 'bestRated'">
        <RecipeList :recipes="bestRatedRecipes" />
      </div>

      <div v-if="activeTab === 'worstRated'">
        <RecipeList :recipes="worstRatedRecipes" />
      </div>
    </div>

    <router-link to="/recipes" class="back-link">Volver al listado</router-link>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import RecipeList from '@/components/RecipeList.vue'
import {type Recipe, useRecipeStore} from '@/stores/recipe'

const activeTab = ref<'bestRated' | 'worstRated'>('bestRated')

const bestRatedRecipes = ref<Recipe[]>([])
const worstRatedRecipes = ref<Recipe[]>([])

const recipeStore = useRecipeStore()

onMounted(async () => {
  const [best, worst] = await Promise.all([
    recipeStore.fetchRecipes(undefined, 'desc', 3),
    recipeStore.fetchRecipes(undefined, 'asc', 3),
  ])

  bestRatedRecipes.value = best
  worstRatedRecipes.value = worst
})

</script>


<style scoped>
h1 {
  text-align: center;
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

.tabs {
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
}

.tabs button {
  padding: 0.5rem 1.5rem;
  border: 1px solid #999;
  background: #eee;
  cursor: pointer;
  border-bottom: none;
  border-radius: 0.5rem 0.5rem 0 0;
  margin: 0 0.25rem;
}

.tabs button.active {
  background: white;
  font-weight: bold;
}

.tab-content {
  border: 1px solid #999;
  border-radius: 0 0 0.5rem 0.5rem;
  padding: 1rem;
}

input[type="number"] {
  width: 3rem;
  text-align: center;
  margin: 0 0.5rem;
}

.back-link {
  display: inline-block;
  margin-bottom: 1rem;
  color: #007bff;
  text-decoration: none;
  font-size: 0.95rem;
}

.back-link:hover {
  text-decoration: underline;
}

select {
  padding: 0.4rem;
  font-size: 1rem;
}

button {
  padding: 0.4rem 1rem;
  font-size: 1rem;
  background-color: #ddd;
  border: 1px solid #999;
  cursor: pointer;
}


</style>
