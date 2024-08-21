<script setup lang="ts">
definePageMeta({ layout: false });
const user = useSupabaseUser();

watch(
  user,
  async () => {
    if (user.value) {
      navigateTo("/");
    }
  },
  { immediate: true },
);
</script>

<template>
  <div>verifying login...</div>
</template>
