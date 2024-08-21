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
        const { username, avatarUrl } = useProfile();
        username.value = data.username;
        avatarUrl.value = data.avatar_url || "";
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
