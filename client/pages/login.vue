<script setup lang="ts">
const email = ref("");
const password = ref("");
</script>

<template>
  <input v-model="email" type="text" placeholder="email" class="bg-inherit" />
  <input
    v-model="email"
    type="text"
    placeholder="password"
    class="bg-inherit"
  />
  <button>Sign up</button>
</template>
