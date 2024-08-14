<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const timeline = response as Post[];
</script>

<template>
  <div class="mb-6 flex items-center justify-between">
    <span class="text-2xl font-bold">Home</span>
    <div class="flex">
      <NuxtLink
        to="/home"
        class="text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home' }"
        >Recent</NuxtLink
      >
      <NuxtLink
        to="/home?s=popular"
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?s=popular' }"
        >Popular</NuxtLink
      >
      <NuxtLink
        to="/home?s=following"
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?s=following' }"
        >Following</NuxtLink
      >
    </div>
  </div>
  <div>
    <Post v-for="post in timeline" v-bind="post" class="mb-2" />
  </div>
</template>
