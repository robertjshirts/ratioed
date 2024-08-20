<script setup lang="ts">
const user = useSupabaseUser();

watch(
  user,
  async () => {
    if (!user.value) return navigateTo("/login");

    return navigateTo("/");
  },
  { immediate: true },
);
</script>

<template>
  <div>Waiting for login...</div>
</template>
