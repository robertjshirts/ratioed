<script setup lang="ts">
const profile = useProfileStore();
</script>

<template>
  <div class="fixed top-0 z-50 w-full border-b bg-inherit backdrop-blur">
    <div class="mx-auto flex max-w-screen-xl justify-between py-3">
      <span @click="navigateTo('/')" class="cursor-pointer text-3xl font-bold"
        >ratioed</span
      >
      <div class="flex items-center">
        <div v-if="profile.username" class="flex items-center">
          <img
            :src="profile.avatarUrl || ''"
            alt=""
            class="me-2 w-8 rounded-full"
          />
          <span>{{ profile.username }}</span>
          <button class="ms-4 flex">
            <Icon name="ph:dots-three-outline-vertical-fill" class="text-xl" />
          </button>
        </div>
        <button
          v-else
          @click="navigateTo('/login')"
          class="rounded-md bg-white px-6 py-1 text-black transition-all hover:bg-gray-200"
        >
          Log in
        </button>
      </div>
    </div>
  </div>
</template>
