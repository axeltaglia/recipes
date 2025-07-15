<template>
  <section v-if="recipe">
    <h1>{{ recipe.name }} ({{ recipe.score }})</h1>

    <div class="tabs">
      <button
        :class="{ active: activeTab === 'ingredientes' }"
        @click="activeTab = 'ingredientes'"
      >
        Ingredientes
      </button>
      <button
        :class="{ active: activeTab === 'pasos' }"
        @click="activeTab = 'pasos'"
      >
        Pasos
      </button>
    </div>

    <div class="tab-content">
      <div v-if="activeTab === 'ingredientes'">
        <label>
          Para
          <input type="number" min="1" v-model.number="people" />
          personas
        </label>
        <ul>
          <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
            {{ ingredient.quantityPerPerson * people }} {{ ingredient.unit }} {{ ingredient.name }}
          </li>
        </ul>
      </div>

      <div v-if="activeTab === 'pasos'">
        <ul>
          <li v-for="(step, index) in recipe.steps" :key="index">
            {{ step }}
          </li>
        </ul>
      </div>
    </div>
  </section>

  <p v-else>Cargando receta...</p>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useRecipeStore } from '@/stores/recipe'

const route = useRoute()
const recipeStore = useRecipeStore()

const recipeId = route.params.id as string
const recipe = ref()

const activeTab = ref<'ingredientes' | 'pasos'>('ingredientes')
const people = ref(4)

onMounted(() => {
  recipe.value = recipeStore.getRecipeById(recipeId)
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
</style>
