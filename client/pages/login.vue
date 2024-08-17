<script setup lang="ts">
const supabase = useSupabaseClient();

definePageMeta({ layout: false });

const loading = ref(false);
const email = ref("");
const password = ref("");
const errorOccured = ref();

const signInWithOtp = async () => {
  const { error } = await supabase.auth.signInWithOtp({
    email: email.value,
    options: {
      emailRedirectTo: "http://localhost:3000/confirm",
    },
  });
  if (error) console.log(error);
};
</script>

<template>
  <div class="flex h-screen w-screen flex-col items-center justify-center">
    <div v-if="errorOccured">{{ errorOccured }}</div>
    <div class="p-3">
      <input
        v-model="email"
        type="text"
        placeholder="email"
        class="rounded-md border bg-inherit px-3 py-2"
      /><br />
      <input
        v-model="password"
        type="text"
        placeholder="password"
        class="rounded-md border bg-inherit px-3 py-2"
      /><br />
    </div>
    <button @click="signInWithOtp" class="flex items-center text-lg">
      <Icon name="ph:envelope-simple"></Icon> Sign in with email
    </button>
  </div>
</template>
