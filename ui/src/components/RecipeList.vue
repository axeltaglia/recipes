<template>
  <table>
    <thead>
    <tr>
      <th>Receta</th>
      <th>Puntaje</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="recipe in recipes" :key="recipe.name">
      <td>{{ recipe.name }}</td>
      <td>{{ recipe.score }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script setup lang="ts">
defineProps<{
  recipes: { name: string; score: number }[]
}>()
</script>

<style scoped>
table {
  border-collapse: collapse;
  width: 100%;
}
th, td {
  padding: 0.5rem 1rem;
  border: 1px solid #ccc;
}
th {
  background-color: #333333;
  text-align: left;
}
</style>
