<script setup lang="ts">
import type { Database } from "~/types/database";
const supabase = useSupabaseClient<Database>();
const profile = useProfileStore();

const loading = ref(false);
const showModal = ref(false);
const content = ref("");
const src = ref("");
const files = ref();

function closeModal() {
  content.value = "";
  src.value = "";
  files.value = null;
  showModal.value = false;
}

function handleFileAttachment(event: any) {
  files.value = event.target.files;
  src.value = URL.createObjectURL(files.value[0]);
}

async function makePost() {
  loading.value = true;
  let attachmentUrl = "";
  if (files.value) {
    const file = files.value[0];
    const fileExt = file.name.split(".").pop();
    const fileName = `${Math.random()}.${fileExt}`;

    const { data: fullPath, error } = await supabase.storage
      .from("attachments")
      .upload(fileName, file);

    attachmentUrl = `https://hhtvcfarcfoxzylxjfij.supabase.co/storage/v1/object/public/${fullPath?.fullPath}`;
  }

  const { error } = await supabase.from("posts").insert({
    profile_id: profile.id,
    content: content.value,
    attachment_url: attachmentUrl,
  });

  closeModal();
  loading.value = false;
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
    class="fixed left-0 top-0 z-50 flex h-screen w-screen items-center justify-center bg-[#0f0f0f] bg-opacity-90 backdrop-blur"
  >
    <div class="flex flex-col rounded-lg border bg-[#131313] p-3 text-white">
      <div v-if="!loading">
        <Icon
          @click="closeModal"
          name="ph:x-circle"
          class="cursor-pointer text-2xl"
        />
        <div class="mt-8 flex border-b px-3">
          <div class="mb-8 flex">
            <img
              :src="profile.avatarUrl || ''"
              class="h-12 w-12 rounded-full"
              alt="pfp failed to load"
            />
            <div class="ms-4 mt-2 flex flex-col">
              <textarea
                v-model="content"
                class="mb-4 w-96 resize-none border-none bg-inherit text-xl outline-none"
                placeholder="What is happening?"
              />
              <div v-if="src" class="max-w-96">
                <button @click="src = ''" class="absolute justify-end p-2">
                  <Icon name="ph:x-circle" class="bg-[#0c1014] text-xl" />
                </button>
                <img
                  :src="src"
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
  </div>
</template>
