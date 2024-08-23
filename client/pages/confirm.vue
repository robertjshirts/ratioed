<script setup lang="ts">
definePageMeta({ layout: false });
const user = useSupabaseUser();
const { username, signIn } = useProfileStore();

watch(
  user,
  async () => {
    if (user.value && !username) {
      await signIn();
      navigateTo("/");
    }
  },
  { immediate: true },
);
</script>

<template>
  <div class="flex h-screen w-screen items-center justify-center">
    <img src="/public/spinner.svg" alt="(loading spinner here)" class="w-12" />
  </div>
</template>
