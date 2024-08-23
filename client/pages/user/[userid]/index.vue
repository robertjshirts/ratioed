<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();
const route = useRoute();
</script>

<template>
  <p>home of {{ route.params.userid }}</p>
</template>
