<script setup lang="ts">
import type { Database } from "~/types/database";
import { RealtimeChannel } from "@supabase/supabase-js";
const props = defineProps<Database["public"]["Views"]["posts_view"]["Row"]>();
const profile = useProfileStore();
const supabase = useSupabaseClient<Database>();

// Post likes or dislikes
const { like, dislike, myReaction } = await useReaction(profile.id ?? null, props.post_id);

// Listen for likes or dislikes
const likeCount = ref(props.likes ?? 0); 
const dislikeCount = ref(props.dislikes ?? 0);

// Ratio posts
const ratioed = useRatio(likeCount, dislikeCount);

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
})

onUnmounted(() => {
  supabase.removeChannel(reactionChannel);
});

</script>

<template>
  <div
    @click="navigateTo(`/user/${user_id}/posts/${post_id}`)"
    class="my-4 flex flex-1 cursor-pointer rounded-lg border bg-[#131313] p-4 pe-8 transition-all hover:bg-[#141414]"
  >
    <div class="mt-1">
      <img
        :src="avatar_url ?? undefined"
        alt=""
        class="min-w-12 max-w-12 rounded-full"
      />
    </div>
    <div class="ms-3 flex flex-1 flex-col">
      <div v-if="parent_id" class="mt-1 text-sm text-gray-400 hover:underline">
        <span @click.stop="navigateTo(`/user/${user_id}/posts/${parent_id}`)">
          Replying to Parent Post
        </span>
      </div>
      <span
        @click="navigateTo(`/user/${user_id}`)"
        class="flex cursor-pointer items-center text-lg font-bold hover:underline"
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
      <span class="mt-1 text-gray-300">{{ ratioed ? "[this post has been ratioed]" : content }}</span>
      <NuxtImg
        preload
        placeholder
        v-if="attachment_url && !ratioed"
        :src="attachment_url"
        alt="image failed to load"
        class="mt-4 cursor-auto rounded-lg bg-contain"
      />
      <div class="mt-4 flex text-gray-400">
        <button
          @click.stop="like"
          class="flex items-center rounded p-2 hover:bg-gray-800"
        >
          <Icon :name="myReaction === 'like' ? 'ph:thumbs-up-fill' : 'ph-thumbs-up'" class="z-10 me-2" />{{ likeCount }}
        </button>
        <button
          @click.stop="dislike"
          class="ms-8 flex items-center rounded p-2 hover:bg-gray-800"
        >
          <Icon :name="myReaction === 'dislike' ? 'ph:thumbs-down-fill' : 'ph:thumbs-down'" class="z-10 me-2" />{{ dislikeCount }}
        </button>
        <button
          @click="navigateTo(`/user/${user_id}/posts/${post_id}`)"
          class="ms-8 flex items-center rounded p-2 hover:bg-gray-800"
        >
          <Icon name="ph:chat" class="me-2" />{{ child_posts }}
        </button>
      </div>
    </div>
  </div>
</template>
