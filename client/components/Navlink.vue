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
    class="flex w-full items-center rounded-lg py-2 text-xl font-bold text-gray-400 transition-all hover:text-gray-200"
    :class="{ 'text-white': $route.path == props.to }"
  >
    <Icon
      :name="$route.path == props.to ? activeIcon : inactiveIcon"
      class="me-4"
    />
    <slot />
  </NuxtLink>
</template>
