import type { Database } from "~/types/database";

export async function useReaction(profile_id: string | null, post_id: string | null) {
    const supabase = useSupabaseClient<Database>();

    const like = async () => {
        if (!profile_id || !post_id) {
            console.error("Profile ID and Post ID are required to like a post");
            return;
        }

        const { error } = await supabase
            .from("reactions")
            .upsert({ profile_id, post_id, reaction_type: "like" })
            .select()
        if (error) {
            console.error("Error liking post", error);
        }
    };

    const dislike = async () => {
        if (!profile_id || !post_id) {
            console.error("Profile ID and Post ID are required to like a post");
            return;
        }
    
        const { data, error } = await supabase
            .from("reactions")
            .select()
            .eq('profile_id', profile_id)
            .eq('post_id', post_id)
            .single();
    
        if (error && error.code !== 'PGRST116') { // 116 means no row found
            console.error("Error checking existing reaction", error);
            return;
        }
    
        if (data) {
            // Update existing reaction
            const { error: updateError } = await supabase
                .from("reactions")
                .update({ reaction_type: "dislike" })
                .eq('profile_id', profile_id)
                .eq('post_id', post_id);
    
            if (updateError) {
                console.error("Error updating dislike", updateError);
            }
        } else {
            // Insert new reaction
            const { error: insertError } = await supabase
                .from("reactions")
                .insert({ profile_id, post_id, reaction_type: "dislike" });
    
            if (insertError) {
                console.error("Error inserting dislike", insertError);
            }
        }
    };

    // const dislike = async () => {
        // if (!profile_id || !post_id) {
            // console.error("Profile ID and Post ID are required to like a post");
            // return;
        // }

        // const { error } = await supabase
            // .from("reactions")
            // .update({ reaction_type: "dislike" })
            // .eq("profile_id", profile_id)
            // .eq('post_id', post_id)
        // if (error) {
            // console.error("Error disliking post", error);
        // }
    // };
    return { like, dislike } as const;
}