<script setup lang="ts">
defineProps<{
  to: string;
  icon: string;
}>();
</script>

<template>
  <NuxtLink
    :to="to"
    class="hover: flex items-center rounded-lg px-4 py-2 text-lg font-bold transition-all"
    :class="{ 'bg-white': $route.path == to, 'text-black': $route.path == to }"
  >
    <Icon :name="icon" class="me-2" />
    <slot></slot>
  </NuxtLink>
</template>
