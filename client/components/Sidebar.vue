<script setup lang="ts">
const supabase = useSupabaseClient();
const user = useSupabaseUser();
const topborder = computed(() => {
  return user;
});
</script>

<template>
  <div class="fixed top-16 box-border h-full w-60">
    <div v-if="user" class="flex flex-col py-8 pl-2">
      <img
        :src="user.user_metadata.avatar_url"
        alt=""
        class="w-24 rounded-full"
      />
      <div class="flex flex-col">
        <span class="mt-4 text-lg">{{ user.user_metadata.full_name }}</span>
        <span class="mt-2 text-lg text-gray-400">@metalcoreenjoyer</span>
      </div>
    </div>
    <nav class="border-b py-5">
      <Navlink to="/home" active-icon="ph:house-fill" inactive-icon="ph:house"
        >Home
      </Navlink>
      <Navlink
        to="/following"
        active-icon="ph:users-three-fill"
        inactive-icon="ph:users-three"
        >Following
      </Navlink>
      <Navlink to="/profile" active-icon="ph:user-fill" inactive-icon="ph:user"
        >My Profile
      </Navlink>
      <Navlink to="/settings" active-icon="ph:gear-fill" inactive-icon="ph:gear"
        >Settings
      </Navlink>
    </nav>
    <div v-if="!user" class="flex flex-col py-3 text-lg">
      <span class="text-gray-300"
        >log in to follow users, ratio others, and have bad takes</span
      >
      <button
        @click="$router.push('/login')"
        class="mb-2 mt-4 rounded-md border py-2"
      >
        Log in
      </button>
    </div>
    <div v-if="user" class="pt-8">
      <PostModal />
    </div>
  </div>
</template>
