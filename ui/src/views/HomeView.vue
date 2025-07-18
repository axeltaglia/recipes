<template>
  <main>
    <h1>Welcome to the Recipe Admin</h1>

    <router-link to="/recipes">Go to Recipes List</router-link>
  </main>
</template>

<style scoped>
main {
  padding: 2rem;
  font-family: sans-serif;
}
h1 {
  color: #333;
}
router-link {
  display: inline-block;
  margin-top: 1rem;
  text-decoration: none;
  color: #007bff;
}
router-link:hover {
  text-decoration: underline;
}
</style>
