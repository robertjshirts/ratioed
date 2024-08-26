<script setup lang="ts">
import { ref } from "vue";
import { useAutoResize } from "~/composables/useAutoResize";
import type { Database } from "~/types/database";

const { postId } = defineProps<{ postId: string }>();
const emit = defineEmits(["reply"]);

const supabase = useSupabaseClient<Database>();
const profile = useProfileStore();

const { content, autoResize } = useAutoResize();
const src = ref("");
const files = ref();

function handleFileAttachment(event: any) {
  files.value = event.target.files;
  src.value = URL.createObjectURL(files.value[0]);
}

async function makeReply() {
  let attachmentUrl = "";
  if (files.value) {
    const file = files.value[0];
    const fileExt = file.name.split(".").pop();
    const fileName = `${Math.random()}.${fileExt}`;

    console.log(`uploading ${file.name} with path of ${fileName}`);

    const { data: fullPath, error } = await supabase.storage
      .from("attachments")
      .upload(fileName, file);

    attachmentUrl = `https://hhtvcfarcfoxzylxjfij.supabase.co/storage/v1/object/public/${fullPath?.fullPath}`;
  }

  const { error } = await supabase.from("posts").insert({
    profile_id: profile.id,
    parent_id: postId,
    content: content.value,
    attachment_url: attachmentUrl,
  });

  if (error) {
    console.error(error);
  }

  content.value = "";
  src.value = "";
  files.value = null;
  emit("reply");
}
</script>

<template>
  <div
    v-if="profile.username"
    class="reply-section mt-4 rounded-lg border bg-[#131313] p-4"
  >
    <div class="flex border-b px-3">
      <img
        :src="profile.avatarUrl ?? undefined"
        class="h-12 w-12 rounded-full"
        alt="pfp failed to load"
      />
      <div class="ms-4 mt-2 w-auto flex-1 flex-col">
        <textarea
          v-model="content"
          class="w-full resize-none border-none bg-inherit text-xl outline-none"
          placeholder="What do you think?"
          @input="autoResize"
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
        @click="makeReply"
        class="rounded-3xl bg-gray-50 px-5 py-2 text-sm font-bold text-black"
      >
        Reply
      </button>
    </div>
  </div>
</template>
