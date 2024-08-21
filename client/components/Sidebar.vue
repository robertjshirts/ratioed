<script setup lang="ts">
const { username, avatarUrl, signOut } = useProfile();
const user = useSupabaseUser();
</script>

<template>
  <div class="fixed top-16 box-border h-full w-60">
    <div v-if="username" class="flex flex-col border-b py-8 pl-2">
      <img
        :src="avatarUrl"
        alt="avatar failed to load"
        class="w-24 rounded-full"
      />
      <div class="flex flex-col">
        <span class="mt-4 text-lg">{{ username }}</span>
        <span class="text-gray-400">{{ user.email }}</span>
      </div>
    </div>
    <nav class="border-b border-[#3f3f3f] py-5">
      <Navlink to="/" active-icon="ph:house-fill" inactive-icon="ph:house"
        >Home
      </Navlink>
      <Navlink
        :to="`/user/user`"
        active-icon="ph:user-fill"
        inactive-icon="ph:user"
        >My Profile
      </Navlink>
      <Navlink to="/settings" active-icon="ph:gear-fill" inactive-icon="ph:gear"
        >Settings
      </Navlink>
      <button
        v-if="username"
        @click="signOut"
        class="flex w-full items-center rounded-lg p-2 text-xl font-bold text-gray-400 transition-all hover:text-gray-200"
      >
        <Icon name="ph:arrow-fat-line-left" class="me-4" />
        Log out
      </button>
    </nav>
    <div v-if="username" class="pt-8">
      <PostModal />
    </div>
    <div v-else class="flex flex-col py-3 text-lg">
      <span class="text-gray-300"
        >log in to follow users, ratio others, and have bad takes</span
      >
      <button
        @click="navigateTo('/login')"
        class="mb-2 mt-4 rounded-md border py-2"
      >
        Log in
      </button>
    </div>
  </div>
</template>
