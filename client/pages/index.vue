<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();

const { data: timeline, status } = await useLazyAsyncData(
  "timeline",
  async () => {
    const { data } = await supabase
      .from("posts_view")
      .select(`*`)
      .is("parent_id", null)
      .order("created_at", { ascending: false });
    return data;
  },
);
</script>

<template>
  <div class="container mb-2 flex items-end justify-between">
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
</template>
