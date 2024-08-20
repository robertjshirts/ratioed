<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();
const user = useSupabaseUser();

const showModal = ref(false);
const content = ref("");
const attachment = reactive({
  file: "",
  url: "",
});

function closeModal() {
  showModal.value = false;
  attachment.url = "";
}

function handleFileAttachment(event: any) {
  const files = event.target.files;
}

async function makePost() {
  const { error } = await supabase.from("posts").insert({
    profile_id: user.value.id,
    content: content.value,
  });
  closeModal();
}
</script>

<template>
  <button
    @click="showModal = true"
    class="w-full rounded-3xl bg-white py-3 text-black transition-all hover:bg-gray-200"
  >
    Post
  </button>
  <div
    v-if="showModal"
    class="fixed left-0 top-0 z-50 flex h-screen w-screen items-center justify-center bg-[#0f0f0f]/80 bg-black bg-opacity-10 backdrop-blur"
  >
    <div class="flex flex-col rounded-lg border bg-[#181818] p-3 text-white">
      <Icon
        @click="closeModal"
        name="ph:x-circle"
        class="cursor-pointer text-2xl"
      />
      <div class="mt-8 flex border-b px-3">
        <div class="mb-8 flex">
          <img
            :src="user.user_metadata.avatar_url"
            class="h-12 w-12 rounded-full"
            alt="pfp failed to load"
          />
          <div class="ms-4 mt-2 flex flex-col">
            <textarea
              v-model="content"
              class="mb-4 w-96 resize-none border-none bg-inherit text-xl outline-none"
              placeholder="What is happening?"
            />
            <div v-if="attachment.url" class="max-w-96">
              <button
                @click="attachment.url = ''"
                class="absolute justify-end p-2"
              >
                <Icon name="ph:x-circle" class="bg-[#0c1014] text-xl" />
              </button>
              <img
                :src="attachment.url"
                alt="attachment failed to load"
                class="rounded-xl bg-cover"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="ms-3 mt-3 flex items-center justify-between">
        <div class="relative">
          <input
            @change="(event: any) => handleFileAttachment(event)"
            type="file"
            accept="image/*"
            class="absolute inset-0 z-10 h-full w-full cursor-pointer opacity-0"
          />
          <div class="flex items-center">
            <Icon name="ph:image" class="text-xl" />
            <span class="ms-2">attach image</span>
          </div>
        </div>
        <button
          @click="makePost"
          class="rounded-3xl bg-white px-5 py-2 text-sm font-bold text-black"
        >
          Post
        </button>
      </div>
    </div>
  </div>
</template>
