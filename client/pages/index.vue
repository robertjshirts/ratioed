<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();

const { data: timeline, status } = await useLazyAsyncData(
  "timeline",
  async () => {
    const { data } = await supabase
      .from("parent_posts_view")
      .select(`*`)
      .order("created_at", { ascending: false });
    return data;
  },
);
</script>

<template>
  <div class="mb-2 flex items-center justify-between">
    <span class="text-2xl font-bold">Home</span>
    <div class="flex">
      <NuxtLink
        to="/"
        class="text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/' }"
        >Recent
      </NuxtLink>
      <NuxtLink
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=popular' }"
        >Popular
      </NuxtLink>
      <NuxtLink
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=following' }"
        >following
      </NuxtLink>
    </div>
  </div>
  <div v-if="status === 'success'">
    <Post v-for="post in timeline" v-bind="post" />
  </div>
  <div v-else>loading posts...</div>
</template>
