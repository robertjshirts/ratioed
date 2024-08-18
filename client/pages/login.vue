<script setup lang="ts">
const supabase = useSupabaseClient();

definePageMeta({ layout: false });

const loading = ref(false);
const email = ref("");
const password = ref("");
const errorOccured = ref();

async function signInWithOtp() {
  const { error } = await supabase.auth.signInWithOtp({
    email: email.value,
    options: {
      emailRedirectTo: "http://localhost:3000/confirm",
    },
  });
  errorOccured.value = error;
}
async function signInWithGoogle() {
  const { error } = await supabase.auth.signInWithOAuth({
    provider: "google",
    options: {
      redirectTo: `http://localhost:3000/confirm`,
    },
  });
  errorOccured.value = error;
}
</script>

<template>
  <div class="flex h-screen w-screen flex-col items-center justify-center">
    <div class="flex flex-col rounded-lg border p-5">
      <div class="mb-4 flex flex-col">
        <span class="text-3xl">Sign in</span>
        <span>Sign in with just an email</span>
      </div>
      <div class="mb-10 mt-4">
        <input
          v-model="email"
          type="text"
          placeholder="email"
          class="w-full rounded-md border bg-inherit px-3 py-2 text-lg"
        /><br />
      </div>
      <div class="flex items-center justify-center">
        <button
          @click="signInWithOtp"
          class="w-full rounded-3xl bg-white p-3 text-center text-black"
        >
          <Icon name="ph:envelope-simple" class="me-2" />Sign in with email
        </button>
      </div>
      <div class="my-5 text-center">-- or --</div>
      <button
        @click="signInWithGoogle"
        class="flex items-center justify-center"
      >
        <Icon name="ph:google-logo-fill" class="me-2 text-3xl" />Sign in with
        google
      </button>
      <div v-if="errorOccured" class="rounded-lg bg-red-500 p-1 text-white">
        {{ errorOccured }}
      </div>
    </div>
  </div>
</template>
