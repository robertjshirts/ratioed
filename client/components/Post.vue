<script setup lang="ts">
import type { Database } from "~/types/database";
defineProps<Database["public"]["Views"]["parent_posts_view"]["Row"]>();
</script>

<template>
  <div class="my-4 flex flex-1 rounded-lg border px-5 py-3">
    <div class="mt-1">
      <img
        :src="avatar_url ?? undefined"
        alt=""
        class="min-w-12 max-w-12 rounded-full"
      />
    </div>
    <div class="ms-3 flex flex-1 flex-col">
      <span class="text-lg font-bold hover:underline">{{ username }}</span>
      <span class="mt-1 text-gray-300">{{ content }}</span>
      <NuxtImg
        preload
        v-if="attachment_url"
        :src="attachment_url"
        alt="image failed to load"
        class="mt-4 rounded-lg bg-contain"
      />
    </div>
  </div>
</template>
