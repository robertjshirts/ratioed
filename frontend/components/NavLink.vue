<script setup lang="ts">
defineProps<{
  to: string;
  icon: string;
}>();
</script>

<template>
  <NuxtLink
    to="/home"
    class="flex items-center rounded-2xl bg-white px-4 py-2 text-lg font-bold text-black"
  >
    <Icon :name="icon" class="me-2" />
    <slot></slot>
  </NuxtLink>
</template>
