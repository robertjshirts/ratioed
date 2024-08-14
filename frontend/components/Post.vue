<script setup lang="ts">
import type { Post } from "~/types/global";

defineProps<Post>();
</script>

<template>
  <article
    @click="$router.push(`/${username}/posts/${id}`)"
    class="flex w-auto cursor-pointer rounded-lg bg-[#1a1a1a] p-4 text-sm transition-all hover:bg-[#1c1c1d] md:text-base"
  >
    <img
      :src="pfp"
      class="max-h-12 max-w-12 rounded-full"
      alt="image failed to load"
    />
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
          class="mt-2 w-96 rounded-lg border border-[#272727] bg-cover"
        />
        <div class="mt-4 flex">
          <button class="flex items-center text-green-500">
            <Icon name="carbon:thumbs-up" />
            <p class="ms-2">{{ likes }}</p>
          </button>
          <button class="ms-8 flex items-center text-red-500">
            <Icon name="carbon:thumbs-down" />
            <p class="ms-2">{{ dislikes }}</p>
          </button>
          <button class="ms-8 flex items-center">
            <Icon name="carbon:chat" />
            <p class="ms-2">12</p>
          </button>
        </div>
      </div>
      <div v-else class="m-4 text-center text-gray-400">
        this post has been ratioed
      </div>
    </div>
  </article>
</template>
