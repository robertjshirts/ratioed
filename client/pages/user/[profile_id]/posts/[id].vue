<script setup lang="ts">
import { RealtimeChannel } from "@supabase/supabase-js";
import { useRoute } from "vue-router";
import type { Database } from "~/types/database";
type Post = Database["public"]["Views"]["posts_view"]["Row"];

const supabase = useSupabaseClient<Database>();

const route = useRoute();
const post_id = route.params.id as string;

const sort = ref<"newest" | "oldest" | "popular">("newest");

let channel: RealtimeChannel;

// Fetch the post
const { data: post, status } = await useLazyAsyncData(
  `post:${post_id}`,
  async () => {
    const { data } = await supabase
      .from("posts_view")
      .select('*')
      .eq("post_id", post_id)
      .single();
    return data;
  },
);

// Fetch the comments
const { data: comments, refresh: refreshComments } = await useLazyAsyncData(
  `comments:${post_id}`,
  async () => {
    let query = supabase
      .from("posts_view")
      .select()
      .eq("parent_id", post_id)

    switch (sort.value) {
      case "newest":
        query = query.order("created_at", { ascending: false });
        break;
      case "oldest":
        query = query.order("created_at", { ascending: true });
        break;
      case "popular":
        query = query.order("likes", { ascending: false });
        break;
    }

    const { data } = await query;
    
    return data;
  },
);

watch(sort, () => {
  refreshComments();
});

onMounted (async function() {
  channel = supabase
    .channel("posts")
    .on(
      "postgres_changes",
      { event: "INSERT", schema: "public", table: "posts" },
      async (payload) => {
        if (payload.new.parent_id !== post_id) return;
        console.log("New comment", payload.new);

        const { post, error }= await usePost(payload.new.id);
        if (error) {
          console.error(error);
        }

        if (!post) return;
        console.log(comments.value)
        comments.value?.unshift(post);
        console.log(comments.value)
      },
    );

  channel.subscribe();
});
</script>

<template>
  <div class="mx-auto py-4">
    <div v-if="status !== 'success'" class="py-4 text-center">
      <p>Loading post...</p>
    </div>

    <div v-else-if="post">

      <Post v-bind="post" />

      <!-- Comments Section -->
      <div class="mt-4">
        <div class="flex">
          <h2 class="text-lg font-bold">Comments</h2>
          <!-- Sort dropdown -->
          <select v-model="sort" class="ms-auto rounded-full p-2 bg-[#131313] text-white">
            <option value="newest">Newest</option>
            <option value="oldest">Oldest</option>
            <option value="popular">Most Popular</option>
          </select>
        </div>
        <ReplyBox :postId="post_id" />
        <div 
        v-if="comments === null || comments.length === 0"
        class="py-4 text-center"
        >
          <p>No comments yet. Be the first!</p>
        </div>
        <div v-else>
          <Comment v-for="comment in comments" v-bind="comment" />
        </div>
      </div>
    </div>

    <div v-else class="py-4 text-center">
      <p>Post not found</p>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
}
.text-center {
  text-align: center;
}
</style>
