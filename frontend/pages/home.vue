<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const timeline = response as Post[];
</script>

<template>
  <Post v-for="post in timeline" v-bind="post" class="mt-2" />
</template>
