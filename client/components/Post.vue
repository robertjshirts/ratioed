<script setup lang="ts">
import type { Database } from "~/types/database";
defineProps<Database["public"]["Views"]["posts_view"]["Row"]>();
</script>

<template>
  <div
    @click="navigateTo(`/user/${user_id}/posts/${post_id}`)"
    class="my-4 flex flex-1 cursor-pointer rounded-lg border bg-[#131313] p-4 pe-8 transition-all hover:bg-[#141414]"
  >
    <div class="mt-1">
      <img
        :src="avatar_url ?? undefined"
        alt=""
        class="min-w-12 max-w-12 rounded-full"
      />
    </div>
    <div class="ms-3 flex flex-1 flex-col">
      <div v-if="parent_id" class="mt-1 text-sm text-gray-400 hover:underline">
        <span @click.stop="navigateTo(`/user/${user_id}/posts/${parent_id}`)">
          Replying to Parent Post
        </span>
      </div>
      <span
        @click="navigateTo(`/user/${user_id}`)"
        class="cursor-pointer text-lg font-bold hover:underline"
        >{{ username }}</span
      >
      <span class="mt-1 text-gray-300">{{ content }}</span>
      <NuxtImg
        preload
        placeholder
        v-if="attachment_url"
        :src="attachment_url"
        alt="image failed to load"
        class="mt-4 cursor-auto rounded-lg bg-contain"
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
