import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RecipesView from "@/views/RecipesView.vue";
import RecipeView from '../views/RecipeView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/recipes', component: RecipesView },
  { path: '/recipes/:id', name: 'RecipeDetail', component: RecipeView }
]


const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router


