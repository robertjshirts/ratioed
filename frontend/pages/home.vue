<script setup lang="ts">
import response from "~/assets/mocks/timeline.json";
const posts = response as any[]
</script>

<template>
  <aside class="h-screen sticky top-0"></aside>
  <main>
    <div v-for="post in posts">
      <Post v-bind="post" />
    </div>
  </main>
</template>
