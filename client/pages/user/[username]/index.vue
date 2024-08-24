<script setup lang="ts">
const route = useRoute();
</script>

<template>
  <p>home of {{ route.params.username }}</p>
</template>
