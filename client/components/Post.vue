<script setup lang="ts">
import type { Database } from "~/types/database";
defineProps<Database["public"]["Views"]["parent_posts_view"]["Row"]>();
</script>

<template>
  <div class="flex py-3">
    <div class="mt-2">
      <img :src="avatar_url ?? undefined" alt="" class="w-12 rounded-full" />
    </div>
    <div class="ms-3 flex flex-col">
      <span class="text-lg font-bold hover:underline">{{ username }}</span>
      <span class="mt-1">{{ content }}</span>
    </div>
  </div>
</template>
