<script setup lang="ts">
import type { Database } from "./types/database.types";
const supabase = useSupabaseClient<Database>();
const user = useSupabaseUser();
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
