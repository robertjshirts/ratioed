<script setup lang="ts">
const supabase = useSupabaseClient();

const username = ref("");
const password = ref("");
</script>

<template>
  <div
    class="fixed left-0 top-0 z-50 flex h-screen w-screen flex-auto items-center justify-center bg-[#0f0f0f]/80 bg-black bg-opacity-10 backdrop-blur"
  >
    <div
      class="flex flex-col items-center rounded-lg border border-[#3f3f3f] bg-[#0f0f0f] p-3 text-white"
    >
      <span class="text-4xl">ratioed</span>
      <div>
        <input v-model="username" type="text" class="bg-inherit" /><br />
        <input v-model="password" type="text" class="bg-inherit" /><br />
        <p>{{ username }} {{ password }}</p>
      </div>
    </div>
  </div>
</template>
