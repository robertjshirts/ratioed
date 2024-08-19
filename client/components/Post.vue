<script setup lang="ts">
defineProps<{
  profile_id: string;
  profiles: {
    username: string;
    avatar_url: string;
  };
  content: string;
  attachment_url?: string;
}>();
</script>

<template>
  <div class="flex py-3">
    <div class="mt-2">
      <img :src="profiles.avatar_url" alt="" class="w-12 rounded-full" />
    </div>
    <div class="ms-3 flex flex-col">
      <span class="text-lg font-bold hover:underline">{{
        profiles.username
      }}</span>
      <span class="mt-1">{{ content }}</span>
    </div>
  </div>
</template>
