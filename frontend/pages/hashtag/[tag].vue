<script setup lang="ts">
const hashtag = useRoute().params.tag;
</script>

<template>
  <p>{{ hashtag }}</p>
</template>
