import type { Database } from "~/types/database";

export async function usePost(post_id: string) {
  const supabase = useSupabaseClient<Database>();
  const { data: post, error } = await supabase
    .from("posts_view")
    .select("*")
    .eq("post_id", post_id)
    .single();

  return { post, error };
}
