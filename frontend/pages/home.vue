<script setup lang="ts">
import type { Post } from "~/types/global";
import response from "~/assets/mocks/timeline.json";

const timeline = response as Post[];
</script>

<template>
  <div class="container mx-auto flex">
    <div class="sticky flex w-3/12 flex-col items-center">
      <div class="w-24">
        <img
          src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"
          alt=""
          class="rounded-full bg-contain"
        />
      </div>
      <p class="mt-4 text-lg font-semibold">Generic User</p>
      <div class="mt-4 flex w-full flex-col">
        <NavLink to="/home" icon="carbon:home">Home</NavLink>
        <NavLink to="/home" icon="carbon:home" class="mt-2">Home</NavLink>
        <NavLink to="/settings" icon="carbon:settings" class="mt-2">
          Settings
        </NavLink>
      </div>
    </div>
    <main class="ms-24 w-full">
      <span class="text-2xl font-bold">Home</span>
      <div v-for="post in timeline" class="mt-2">
        <Post v-bind="post" />
      </div>
    </main>
  </div>
</template>
