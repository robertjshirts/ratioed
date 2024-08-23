<script setup lang="ts">
import { useRoute } from "vue-router";
import { usePost } from "~/composables/usePost";
import type { Database } from "~/types/database";

const supabase = useSupabaseClient<Database>();

const route = useRoute();
const post_id = route.params.id as string;

// const { post, error, loading } = await usePost(post_id);


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

const { data: comments, refresh } = await useLazyAsyncData(
  `comments:${post_id}`,
  async () => {
    const { data } = await supabase
      .from("posts_view")
      .select()
      .eq("parent_id", post_id)
      .order("created_at", { ascending: false });
    return data;
  },
);
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
        <h2 class="text-lg font-bold">Comments</h2>
        <ReplyBox :postId="post_id" @reply="refresh()" />
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
