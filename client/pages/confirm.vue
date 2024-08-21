<script setup lang="ts">
const router = useRouter();
const user = useSupabaseUser();

setTimeout(() => router.push("/login"), 5000);

watch(
  user,
  async () => {
    if (user.value) {
      const supabase = useSupabaseClient();

      const { data } = await supabase
        .from("profiles")
        .select(`id, username, avatar_url`)
        .eq("id", user.value.id)
        .single();

      if (data) {
        const { username, email, avatarUrl } = useProfile();
        username.value = data.username;
        email.value = user.value.email || "";
        avatarUrl.value = data.avatar_url || "";
      }
      router.push("/");
    }
  },
  { immediate: true },
);
</script>

<template>
  <div>Waiting for login...</div>
</template>
