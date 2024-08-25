<script setup lang="ts">
import { RealtimeChannel } from "@supabase/supabase-js";
import type { Database } from "~/types/database";
type Post = Database["public"]["Views"]["posts_view"]["Row"];
const supabase = useSupabaseClient<Database>();

let channel: RealtimeChannel;
let timeline = ref<Post[]>([]);

onMounted(async function() {
  const { data, error } = await supabase
    .from("posts_view")
    .select(`*`)
    .is("parent_id", null)
    .order("created_at", { ascending: false });
  timeline.value = data ?? [];
  if (error) {
    console.error(error);
  }

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

onUnmounted(() => {
  supabase.removeChannel(channel);
});

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
  </div>
</template>
