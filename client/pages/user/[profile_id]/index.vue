<script setup lang="ts">
const route = useRoute();
import type { Database } from "~/types/database";
</script>

<template>
  <p>Home of {{ route.params.profile_id }}</p>
</template>
