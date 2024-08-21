<script setup lang="ts">
defineProps<{
  profile_id: string | null;
  profiles: {
    username: string | null;
    avatar_url: string | null;
  };
  content: string | null;
  attachment_url?: string | null;
  parent_id: string | null;
}>();
</script>

<template>
  <div class="ms-4 flex py-2">
    <div class="mt-1">
      <img
        :src="profiles.avatar_url ?? undefined"
        alt="Avatar"
        class="w-8 rounded-full"
      />
    </div>
    <div class="ms-2 flex flex-col">
      <span class="text-sm font-semibold text-gray-700 hover:underline">{{
        profiles.username
      }}</span>
      <span class="mt-1 text-xs text-gray-600">{{ content }}</span>
    </div>
  </div>
</template>
