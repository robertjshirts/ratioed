import type { Database } from "~/types/database";
import { ref } from "vue";

export async function usePost(post_id: string) {
    type ParentPostsView = Database["public"]["Views"]["parent_posts_view"];
    type Post = ParentPostsView["Row"];

    const supabase = useSupabaseClient<Database>();
    
    const post = ref<Post | null>(null);
    const error = ref<string | null>(null);
    const loading = ref<boolean>(true);

    const fetchPost = async () => {
        try {
            const { data, error: fetchError } = await supabase
                .from("parent_posts_view")
                .select("*")
                .eq("post_id", post_id)
                .single();

            if (fetchError) {
                throw fetchError;
            }

            post.value = data;
        } catch (err: any) {
            console.log("Error fetching post", err);
            error.value = err.message || "Unknown error";
        } finally {
            loading.value = false;
        }
    };

    fetchPost();

    return { post, error, loading } as const;
}