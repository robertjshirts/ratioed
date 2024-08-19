<script setup lang="ts">
import type { Database, Tables } from "~/types";
const supabase = useSupabaseClient<Database>();
const user = useSupabaseUser();

const timeline = ref<any[]>([]);

const { data, error } = await supabase
  .from("posts")
  .select(`*, profiles:profile_id (*)`);
if (data) timeline.value = data;
</script>

<template>
  <Sidebar />
  <div class="pl-80 pt-8">
    <div class="mb-2 flex items-center justify-between">
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
    <div>{{ user }}</div>
    <Post v-for="post in timeline" v-bind="post" />
  </div>
</template>
