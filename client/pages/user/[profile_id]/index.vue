<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();
const route = useRoute();

const { data } = await useAsyncData("profile", async () => {
  const { data, error } = await supabase
    .from("profiles")
    .select("*")
    .eq("id", route.params.profile_id)
    .single();

  if (error) {
    return null;
  }

  return data;
});
</script>

<template>
  <div>
    <NuxtImg :src="data!.avatar_url" class="w-32 rounded-full" />
  </div>
</template>
