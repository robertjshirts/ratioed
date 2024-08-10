<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const posts = response as Post[];
</script>

<template>
  <main>
    <div v-for="post in posts">
      <Post v-bind="post" />
    </div>
  </main>
</template>
