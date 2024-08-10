<script setup lang="ts">
import type { Post } from "~/types/global";

defineProps<Post>();
</script>

<template>
  <div
    class="transition-allmd:text-base flex w-auto border p-4 text-sm"
  >
    <img :src="pfp" class="max-w-12 max-h-12 rounded-full"> 
    <div class="ms-4">
      <div class="flex">
        <p class="font-bold">{{ username }}</p>
        <p class="ms-2 text-gray-400">·</p>
        <p class="ms-2 text-gray-400">{{ timestamp }}</p>
      </div>
      <div v-if="!content.ratioed">
        <pre class="mt-2 font-sans">{{ content.body }}</pre>
        <img
          v-if="content.attachment"
          :src="content.attachment.src"
          alt="image failed to load"
          class="mt-2 rounded-xl border border-slate-800"
        />
        <div class="mt-4 flex">
          <button class="flex items-center">
            <Icon name="carbon:thumbs-up" />
            <p class="ms-2">{{ likes }}</p>
          </button>
        </div>
      </div>
      <div v-else class="m-4 text-center text-gray-400">
        this post has been ratioed
      </div>
    </div>
  </div>
</template>
