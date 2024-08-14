<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const timeline = response as Post[];
</script>

<template>
  <span class="text-2xl font-bold">Home</span>
  <div v-for="post in timeline" class="mt-2">
    <Post v-bind="post" />
  </div>
</template>
