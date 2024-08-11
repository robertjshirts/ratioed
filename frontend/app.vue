<script setup lang="ts"></script>

<template>
  <div class="fixed z-40 w-full border bg-white px-4">
    <div class="flex items-center justify-between p-3">
      <span class="text-2xl font-semibold">ratioed</span>
      <div>
        <input type="text" class="rounded-full px-3" />
      </div>
      <button>login</button>
    </div>
  </div>
  <div class="pt-14">
    <NuxtPage />
  </div>
</template>
