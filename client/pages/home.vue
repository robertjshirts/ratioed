<script setup lang="ts">
const supabase = useSupabaseClient();

const posts = ref<any>([]);

supabase
  .channel("posts")
  .on(
    "postgres_changes",
    { event: "INSERT", schema: "public", table: "posts" },
    (payload) => {
      posts.value.push(payload.new);
    },
  )
  .subscribe();
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
  <Post />
</template>
