<script setup lang="ts">
import type { Database } from "~/types/database";
defineProps<Database["public"]["Views"]["parent_posts_view"]["Row"]>();
</script>

<template>
  <div
    class="my-4 flex flex-1 rounded-lg border bg-[#131313] p-4 pe-8 hover:bg-[#141414]"
  >
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
        placeholder
        v-if="attachment_url"
        :src="attachment_url"
        alt="image failed to load"
        class="mt-4 rounded-lg bg-contain"
      />
      <div class="mt-4 flex text-gray-400">
        <button class="flex items-center">
          <Icon name="ph:thumbs-up" class="z-10 me-2" />{{ likes }}
        </button>
        <button class="flex items-center">
          <Icon name="ph:thumbs-down" class="z-10 me-2 ms-8" />{{ dislikes }}
        </button>
        <button class="flex items-center">
          <Icon name="ph:chat" class="me-2 ms-8" />{{ child_posts }}
        </button>
      </div>
    </div>
  </div>
</template>
