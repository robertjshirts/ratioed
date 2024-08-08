<script setup lang="ts">
import * as posts from '~/assets/mocks/timeline.json'
</script>

<template>
  <div class="flex justify-center">
    <aside class="h-screen sticky top-0"></aside>
    <main>
      <div v-for="post in posts">
        <Post :post="post"/>
      </div>
    </main>
  </div>
</template>
