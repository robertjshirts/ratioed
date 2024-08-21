<script setup lang="ts">
import type { Database } from "~/types/database";
defineProps<Database["public"]["Views"]["parent_posts_view"]["Row"]>();
</script>

<template>
  <div class="flex border-t py-3">
    <div class="mt-1">
      <img
        :src="avatar_url ?? undefined"
        alt=""
        class="min-w-12 max-w-12 rounded-full"
      />
    </div>
    <div class="ms-3 flex flex-col">
      <span class="text-lg font-bold hover:underline">{{ username }}</span>
      <span class="mt-1">{{ content }}</span>
      <img
        v-if="attachment_url"
        :src="attachment_url"
        alt=""
        class="mt-4 w-64 rounded-lg bg-contain"
      />
    </div>
  </div>
</template>
