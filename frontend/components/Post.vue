<script setup lang="ts">
defineProps<{
  pfp: string
  username: string
  timestamp: string
  content: {
    ratioed?: boolean
    body?: string
    attachment?: {
      type: string
      src: string
    }
  }
  likes: number
  dislikes: number
}>()
</script>

<template>
  <div
    class="flex p-4 border border-slate-800 w-auto hover:bg-[#1e1d24] transition-all"
  >
    <UAvatar :src="pfp" alt="avatar" size="lg" />
    <div class="ms-4">
      <div class="flex">
        <p class="font-bold">{{ username }}</p>
        <p class="ms-2 text-gray-400">·</p>
        <p class="ms-2 text-gray-400">{{ timestamp }}</p>
      </div>
      <div v-if="content.ratioed">
        <pre class="mt-2 font-sans">{{ content.body }}</pre>
        <img
          v-if="content.attachment"
          :src="content.attachment.src"
          alt="image failed to load"
          class="mt-2 rounded-xl border border-slate-800"
        />
        <div class="flex mt-4">
          <UButton
            icon="carbon:thumbs-up"
            color="green"
            variant="ghost"
            :padded="false"
            >{{ likes }}</UButton
          >
          <UButton
            icon="carbon:thumbs-down"
            color="red"
            variant="ghost"
            :padded="false"
            class="ms-4"
            >{{ dislikes }}</UButton
          >
          <UButton
            icon="carbon:chat"
            color="white"
            variant="ghost"
            :padded="false"
            class="ms-4"
            >{{ likes }}</UButton
          >
        </div>
      </div>
      <div v-else class="text-gray-400 m-4 text-center">
        this post has been ratioed
      </div>
    </div>
  </div>
</template>
