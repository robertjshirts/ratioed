<script setup lang="ts">
definePageMeta({ layout: false });

const supabase = useSupabaseClient();
const url = useRequestURL();

const loading = ref(false);
const username = ref("");
const email = ref("");
const errorOccured = ref();

async function signInWithOtp() {
  loading.value = true;

  const { error } = await supabase.auth.signInWithOtp({
    email: email.value,
    options: {
      data: {
        username: username.value,
      },
      emailRedirectTo: `${url.origin}/confirm`,
    },
  });

  errorOccured.value = error;
  loading.value = false;
}

async function signInWithGoogle() {
  loading.value = true;

  const { data, error } = await supabase.auth.signInWithOAuth({
    provider: "google",
    options: {
      redirectTo: `${url.origin}/confirm`,
    },
  });

  errorOccured.value = error;
  loading.value = false;
}
</script>

<template>
  <div class="flex h-screen w-screen flex-col items-center justify-center">
    <div class="flex max-w-sm flex-col rounded-lg border bg-[#181818] p-7">
      <div class="mb-4 flex flex-col">
        <span class="text-3xl">Sign in</span>
        <span class="text-gray-400"
          >Sign in with one of the following methods</span
        >
      </div>
      <button
        @click="signInWithGoogle"
        class="mt-4 flex items-center justify-center rounded-md border p-3"
      >
        <img src="/public/google-icon.svg" alt="" class="me-6 w-5" />
        <span>Sign in with google</span>
      </button>
      <span class="flex justify-center py-4 text-gray-400">-- or --</span>
      <form class="mb-4 text-gray-300">
        <label>username</label>
        <input
          v-model="username"
          type="text"
          class="mb-3 mt-1 w-full rounded-md border bg-inherit px-3 py-2 text-lg"
        />
        <label>email</label>
        <input
          v-model="email"
          type="text"
          class="mt-1 w-full rounded-md border bg-inherit px-3 py-2 text-lg"
        />
      </form>
      <div class="mt-4 flex items-center justify-center align-middle">
        <button
          @click="signInWithOtp"
          class="w-full rounded-md bg-white p-3 text-center font-bold text-black"
        >
          <Icon name="ph:envelope-simple" class="me-2" />{{
            loading ? "loading" : "sign in with magic link"
          }}
        </button>
      </div>
      <div v-if="errorOccured" class="rounded-lg bg-red-500 p-1 text-white">
        {{ errorOccured }}
      </div>
    </div>
  </div>
</template>
