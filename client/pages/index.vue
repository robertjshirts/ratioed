<script setup lang="ts">
const supabase = useSupabaseClient();
const user = useSupabaseUser();

const timeline = ref<any[]>([]);

// const { data, error } = await supabase
//   .from("posts")
//   .select(`*, profiles:profile_id (*)`);
// if (data) timeline.value = data;
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
  <div>{{ user }}</div>
  <Post v-for="post in timeline" v-bind="post" />
</template>
