<script setup lang="ts">
import type { Database } from "~/types/database";

const supabase = useSupabaseClient<Database>();
const user = useSupabaseUser();
const profile = useProfile();

watch(
  user,
  async () => {
    if (user.value) {
      const { data } = await supabase
        .from("profiles")
        .select(`username, avatar_url`)
        .eq("id", user.value.id)
        .single();

      if (data) {
        profile.value.username = data.username;
        profile.value.email = user.value.email;
        profile.value.avatarUrl = data.avatar_url;
      }

      return navigateTo("/");
    }
  },
  { immediate: true },
);
</script>

<template>
  <div>Waiting for login...</div>
</template>
