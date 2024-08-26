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
  <div class="flex">
    <NuxtImg :src="data?.avatar_url" class="w-28 rounded-full" />
    <div class="ms-2 mt-1">
      <span class="text-2xl">{{ data?.username }}</span
      ><br />
      <span>{{ data?.bio }}</span>
    </div>
  </div>
  <span>page still in construction</span>
</template>
