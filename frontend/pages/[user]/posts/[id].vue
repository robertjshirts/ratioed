<script setup lang="ts">
const user = useRoute().params.user;
const postId = useRoute().params.id;
</script>

<template>
  <p>post with id of {{ postId }} made by {{ user }}</p>
  <NavLink to="home" icon="">home</NavLink>
</template>
