<script setup lang="ts">
import { useRoute } from 'vue-router';
import { usePost } from '~/composables/usePost';
import type { Database } from "~/types/database";
import PostComponent from '~/components/Post.vue';
import CommentComponent from '~/components/Comment.vue';

const supabase = useSupabaseClient();

const route = useRoute();
const post_id = route.params.id as string;

const { post, error, loading } = await usePost(post_id);    

type ParentPostsView = Database["public"]["Views"]["parent_posts_view"];
type Post = ParentPostsView["Row"];

const comments = ref<Post[]>([]);
const commentsError = ref<string | null>(null);

watch(post, async (postValue) => {
  if (postValue) {
    const { data, error } = await supabase
      .from('parent_posts_view')
      .select()
      .eq('parent_id', post_id);
    if (error) {
      console.error(error);
      commentsError.value = error.message;
    }
    if (data) {
      comments.value = data;
    }
  }
});

</script>

<template>
  <div class="container mx-auto py-4">
    <div v-if="loading" class="text-center py-4">
      <p>Loading post...</p>
    </div>

    <div v-else-if="error" class="text-center py-4 text-red-500">
      <p>Error: {{ error }}</p>
    </div>

    <div v-else-if="post">
      <PostComponent
        :profile_id="post.user_id"
        :profiles="{
          username: post.username,
          avatar_url: post.avatar_url,
        }"
        :content="post.content"
        :attachment_url="post.attachment_url"
      />

      <!-- Comments Section -->
      <div class="mt-4">
        <h2 class="text-lg font-bold">Comments</h2>
        <div v-if="commentsError" class="text-red-500">
          <p>Error loading comments: {{ commentsError }}</p>
        </div>
        <div v-else-if="comments.length === 0">
          <p>No comments yet.</p>
        </div>
        <div v-else>
          <ul>
            <li v-for="comment in comments" :key="comment.post_id ?? undefined" class="mt-2">
              <CommentComponent
                :profile_id="comment.user_id"
                :profiles="{
                  username: comment.username,
                  avatar_url: comment.avatar_url,
                }"
                :content="comment.content"
                :attachment_url="comment.attachment_url"
                :parent_id="comment.parent_id"
              />
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-4">
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
