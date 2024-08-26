<script setup lang="ts">
import { RealtimeChannel } from "@supabase/supabase-js";
import type { Database } from "~/types/database";
type Post = Database["public"]["Views"]["posts_view"]["Row"];
const supabase = useSupabaseClient<Database>();

let channel: RealtimeChannel;
let timeline = ref<Post[]>([]);

const limit = 25;
let offset = 0;

// Fetch timeline onMounted
onMounted(async function() {
  const { data, error } = await supabase
    .from("posts_view")
    .select(`*`)
    .is("parent_id", null)
    .order("created_at", { ascending: false })
    .limit(limit);
  timeline.value = data ?? [];
  if (error) {
    console.error(error);
  }

  // Subscribe for realtime posts
  channel = supabase
    .channel("posts")
    .on(
      "postgres_changes",
      { event: "INSERT", schema: "public", table: "posts" },
      async (payload) => {
        const post = await getPostView(payload.new.id);

        if (!post) return;

        timeline.value.unshift(post);
      },
    );

  channel.subscribe();
});

// Unmount the channel
onUnmounted(() => {
  supabase.removeChannel(channel);
});

// Load more posts
async function loadMore() {
  offset += limit;
  const { data, error } = await supabase
    .from("posts_view")
    .select(`*`)
    .is("parent_id", null)
    .order("created_at", { ascending: false })
    .range(offset, offset + limit - 1);
  if (error) {
    console.error(error);
  }

  timeline.value.push(...(data ?? []));
}

// Util function
async function getPostView(post_id: string) {
  const { data, error } = await supabase
    .from("posts_view")
    .select(`*`)
    .eq("post_id", post_id)
    .single();
  if (error) {
    console.error(error);
  }
  return data;
}
</script>

<template>
  <div
    class="mb-2 me-8 ms-8 flex items-end justify-between md:ms-0 lg:me-12 xl:me-72"
  >
    <span class="text-2xl font-bold">Home</span>
    <div class="flex">
      <NuxtLink
        to="/"
        class="text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/' }"
        >Recent
      </NuxtLink>
      <NuxtLink
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=popular' }"
        >Popular
      </NuxtLink>
      <NuxtLink
        class="ms-4 text-gray-400"
        :class="{ 'text-white': $route.fullPath == '/home?f=following' }"
        >Following
      </NuxtLink>
    </div>
  </div>
  <div class="me-8 ms-8 md:ms-0 xl:me-72">
    <Post v-for="post in timeline" :key="post.post_id ?? undefined" v-bind="post" />
    <button
      v-if="timeline.length >= (offset + limit)"
      @click="loadMore"
      class="w-full rounded-full bg-white px-2 py-3 mt-2 mb-10 text-black"
      >
      Load more
    </button>
    <p v-else class="pt-4 mb-10 text-center">
      Damn you read all the posts. Touch grass, dude. 
    </p>
  </div>
</template>
