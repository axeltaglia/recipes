<template>
  <section v-if="recipe">
    <h1>{{ recipe.name }} ({{ recipe.score }})</h1>

    <div class="tabs">
      <button
          :class="{ active: activeTab === 'ingredients' }"
          @click="activeTab = 'ingredients'"
      >
        Ingredientes
      </button>
      <button
          :class="{ active: activeTab === 'steps' }"
          @click="activeTab = 'steps'"
      >
        Pasos
      </button>
    </div>

    <div class="tab-content">
      <div v-if="activeTab === 'ingredients'">
        <label>
          Para
          <input type="number" min="1" v-model.number="people" />
          personas
        </label>
        <ul>
          <li
              v-for="(ri, index) in recipe.ingredients"
              :key="index"
          >
            {{ (ri.quantityPerPerson * people).toFixed(2)  }}
            {{ ri.ingredient.unit }}
            {{ ri.ingredient.name }}
          </li>
        </ul>
      </div>

      <div v-if="activeTab === 'steps'">
        <ul>
          <li
              v-for="step in recipe.steps"
              :key="step.id"
          >
            {{ step.description }}
          </li>
        </ul>
      </div>
    </div>
    <router-link to="/recipes" class="back-link">Volver al listado</router-link>
  </section>

  <p v-else>Cargando receta...</p>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useRecipeStore } from '@/stores/recipe'
import { storeToRefs } from 'pinia'

const route = useRoute()
const recipeStore = useRecipeStore()
const { selectedRecipe } = storeToRefs(recipeStore)

const activeTab = ref<'ingredients' | 'steps'>('ingredients')
const people = ref(4)

onMounted(() => {
  const id = Number(route.params.id)
  recipeStore.fetchRecipeById(id)
})

const recipe = computed(() => selectedRecipe.value)
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

</style>
