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
  <div class="flex py-2 ms-4">
    <div class="mt-1">
      <img :src="profiles.avatar_url ?? undefined" alt="Avatar" class="w-8 rounded-full" />
    </div>
    <div class="ms-2 flex flex-col">
      <span class="text-sm font-semibold hover:underline text-gray-700">{{ profiles.username }}</span>
      <span class="text-xs mt-1 text-gray-600">{{ content }}</span>
    </div>
  </div>
</template>

<style scoped>
.flex {
  align-items: flex-start;
}
.ms-4 {
  margin-left: 1rem;
}
.text-sm {
  font-size: 0.875rem; /* Smaller font for comments */
}
.text-xs {
  font-size: 0.75rem; /* Even smaller for content */
}
.text-gray-700 {
  color: #4a5568; /* Slightly muted text color */
}
.text-gray-600 {
  color: #718096; /* Even more muted for the comment content */
}
.w-8 {
  width: 2rem;
}
</style>
