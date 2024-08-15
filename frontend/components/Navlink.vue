<script setup lang="ts">
const props = defineProps<{
  to: string;
  activeIcon: string;
  inactiveIcon: string;
}>();
</script>

<template>
  <NuxtLink
    :to="to"
    class="hover: flex w-full items-center rounded-lg p-2 text-xl font-bold text-gray-400 transition-all"
    :class="{ 'text-white': $route.path == props.to }"
  >
    <Icon
      :name="$route.path == props.to ? activeIcon : inactiveIcon"
      class="me-4"
    />
    <slot />
  </NuxtLink>
</template>
