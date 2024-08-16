<script setup lang="ts">
const route = useRoute();
</script>

<template>
  <p>home of {{ route.params.user }}</p>
</template>
