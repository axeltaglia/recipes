import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RecipesView from "@/views/RecipesView.vue";

const routes = [
  { path: '/', component: HomeView },
  { path: '/recipes', component: RecipesView }
]


const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router


