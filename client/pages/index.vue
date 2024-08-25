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
  <div
    class="mb-2 me-8 ms-8 flex items-end justify-between md:ms-0 lg:me-12 xl:me-72"
  >
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
  <div v-if="status === 'success'" class="me-8 ms-8 md:ms-0 xl:me-72">
    <Post v-for="post in timeline" v-bind="post" />
  </div>
</template>
