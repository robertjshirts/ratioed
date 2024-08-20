<script setup lang="ts">
import type { Database } from "~/types/database";

const user = useSupabaseUser();

watch(
  user,
  async () => {
    if (user.value) {
      const supabase = useSupabaseClient<Database>();

      const { data } = await supabase
        .from("profiles")
        .select(`id, username, avatar_url`)
        .eq("id", user.value.id)
        .single();

      if (data) {
        const profile = useProfile();
        profile.value.username = data.username;
        profile.value.avatarUrl = data.avatar_url;
      }

      reloadNuxtApp();
      return navigateTo("/");
    }
  },
  { immediate: true },
);
</script>

<template>
  <div>Waiting for login...</div>
</template>
