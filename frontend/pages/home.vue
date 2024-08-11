<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const timeline = response as Post[];
</script>

<template>
  <main>
    <div v-for="post in timeline">
      <Post v-bind="post" />
    </div>
  </main>
</template>
