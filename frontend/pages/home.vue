<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const posts = response as Post[]
</script>

<template>
  <aside class="h-screen sticky top-0"></aside>
  <main>
    <div v-for="post in posts">
      <Post v-bind="post" />
    </div>
  </main>
</template>
