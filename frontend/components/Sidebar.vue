<script setup lang="ts">
const status = true;
</script>

<template>
  <div class="fixed top-16 box-border h-full w-60">
    <div v-if="status" class="flex flex-col py-8">
      <img
        src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSUxlLWqyirlf0ApuqexYrx3jH9P031hUiCOw&s"
        alt=""
        class="w-24 rounded-full"
      />
      <div class="flex flex-col">
        <span class="mt-4 text-lg">Heavener Fan</span>
        <span class="mt-2 text-lg text-gray-400">@metalcoreenjoyer</span>
      </div>
    </div>
    <nav class="border-b border-[#3f3f3f] py-5" :class="{ 'border-y': status }">
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
    <div v-if="!status" class="flex flex-col py-3 text-lg">
      <span class="text-gray-300"
        >log in to follow users, ratio others, and have bad takes</span
      >
      <button class="mb-2 mt-4 rounded-md border py-2">Log in</button>
    </div>
    <div v-if="status" class="pt-8">
      <button class="w-full rounded-3xl bg-white py-3 text-black">Post</button>
    </div>
  </div>
</template>
