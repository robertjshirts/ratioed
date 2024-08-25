<script setup lang="ts">
import type { Database } from "~/types/database";
import { RealtimeChannel } from '@supabase/supabase-js';
const props = defineProps<Database["public"]["Views"]["posts_view"]["Row"]>();
const profile = useProfileStore();
const supabase = useSupabaseClient();

// Post likes or dislikes
const { like, dislike } = await useReaction(profile.id ?? null, props.post_id);

// Listen for likes or dislikes
const likeCount = ref(props.likes ?? 0);
const dislikeCount = ref(props.dislikes ?? 0);

let reactionChannel: RealtimeChannel;
onMounted(() => {
  reactionChannel = supabase
    .channel(`reactions:post_id=eq.${props.post_id}`)
    .on("postgres_changes", 
      {event: "INSERT", schema: "public", table: "reactions"},
      (payload) => {
        if (payload.new.post_id !== props.post_id) return;
        if (payload.new.reaction_type === "like") {
            likeCount.value++;
        } else if (payload.new.reaction_type === "dislike") {
            dislikeCount.value++;
        }
      }
    )
    .on("postgres_changes",
      {event: "UPDATE", schema: "public", table: "reactions"},
      (payload) => {
        if (payload.new.post_id !== props.post_id) return;
        if (payload.errors) {
          console.error(payload.errors);
          return;
        }
        console.log("Someone updated a reaction to: ", payload.new.reaction_type);
        if (payload.new.reaction_type === "like") {
            likeCount.value++;
            dislikeCount.value--;
        } else if (payload.new.reaction_type === "dislike") {
            dislikeCount.value++;
            likeCount.value--;
        }
      }
    );

  reactionChannel.subscribe();
});
</script>

<template>
  <div 
  @click="navigateTo(`/user/${user_id}/posts/${post_id}`)"
  class="my-2 flex flex-1 cursor-pointer rounded-md border bg-[#1a1a1a] p-3 hover:bg-[#202020]"
  >
    <div class="mt-1">
      <img
        :src="avatar_url ?? undefined"
        alt=""
        class="min-w-10 max-w-10 rounded-full"
      />
    </div>
    <div class="ms-2 flex flex-1 flex-col">
      <span 
      @click="navigateTo(`/user/${user_id}`)"
      class="text-base font-semibold hover:underline"
      >{{ username }}
        <Icon
          v-if="role == 'verified'"
          name="ph:circle-wavy-check-duotone"
          class="ms-2 text-blue-500"
        />
        <Icon
          v-if="role == 'dev'"
          name="ph:code-bold"
          class="ms-2 text-green-500"
        />
    </span>

      <span class="mt-1 text-gray-400">{{ content }}</span>
      <NuxtImg
        preload
        placeholder
        v-if="attachment_url"
        :src="attachment_url"
        alt="image failed to load"
        class="mt-3 rounded-md bg-contain"
      />
      <div class="mt-3 flex text-gray-500 text-sm">
        <button @click="like" class="flex items-center">
          <Icon name="ph:thumbs-up" class="z-10 me-2" />{{ likes }}
        </button>
        <button @click="dislike" class="flex items-center">
          <Icon name="ph:thumbs-down" class="z-10 me-4 ms-6" />{{ dislikes }}
        </button>
        <button @click="navigateTo(`/user/${user_id}/posts/${post_id}`)" class="flex items-center p-2 hover:bg-gray-800 rounded ms-8">
          <Icon name="ph:chat" class="me-2" />{{ child_posts }}
        </button>
      </div>
    </div>
  </div>
</template>
