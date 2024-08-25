<script setup lang="ts">
import type { Database } from "~/types/database";
const props = defineProps<Database["public"]["Views"]["posts_view"]["Row"]>();
const profile = useProfileStore();
const { like, dislike } = await useReaction(profile.id ?? null, props.post_id);
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
        <button @click="like" class="flex items-center p-2 hover:bg-gray-800 rounded">
          <Icon name="ph:thumbs-up" class="z-10 me-2" />{{ likes }}
        </button>
        <button @click="dislike" class="flex items-center p-2 hover:bg-gray-800 rounded ms-8">
          <Icon name="ph:thumbs-down" class="z-10 me-2" />{{ dislikes }}
        </button>
        <button @click="navigateTo(`/user/${user_id}/posts/${post_id}`)" class="flex items-center p-2 hover:bg-gray-800 rounded ms-8">
          <Icon name="ph:chat" class="me-2" />{{ child_posts }}
        </button>
      </div>
    </div>
  </div>
</template>
