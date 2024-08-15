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
        >Recent
      </NuxtLink>
      <NuxtLink
        to="/home?f=popular"
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=popular' }"
        >Popular
      </NuxtLink>
      <NuxtLink
        to="/home?f=following"
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=following' }"
        >following
      </NuxtLink>
    </div>
  </div>
  <div>
    <Post v-for="post in timeline" v-bind="post" class="mb-2" />
  </div>
</template>
