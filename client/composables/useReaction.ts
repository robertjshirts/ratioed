import { ref } from "vue";
import { RealtimeChannel } from "@supabase/supabase-js";
import type { Database } from "~/types/database";

export async function useReaction(profile_id: string | null, post_id: string | null) {
    const supabase = useSupabaseClient<Database>();
    const myReaction = ref<"like" | "dislike" | null>(null);

    const like = async () => {
        if (!profile_id || !post_id) {
            console.log("User attempted to react to a post without being logged in");
            return;
        }

        if (myReaction.value === "like") {
            return;
        }

        const { error } = await supabase
            .from("reactions")
            .upsert({ profile_id, post_id, reaction_type: "like" })
            .select()
        if (error) {
            console.error("Error liking post", error);
        }
        myReaction.value = "like";
    };

    const dislike = async () => {
        if (!profile_id || !post_id) {
            console.log("User attempted to react to a post without being logged in");
            return;
        }

        if (myReaction.value === "dislike") {
            return;
        }

        const { error } = await supabase
            .from("reactions")
            .upsert({ profile_id, post_id, reaction_type: "dislike" })
            .select()
        if (error) {
            console.error("Error disliking post", error);
        }
        myReaction.value = "dislike";
    };

    if (!profile_id || !post_id) {
        return { like, dislike, myReaction } as const;
    }

    // Check if user has already reacted to this post
    const { data: myReactionData } = await supabase
        .from("reactions")
        .select("reaction_type")
        .eq("profile_id", profile_id)
        .eq("post_id", post_id)
        .single();
    if (myReactionData) {
        myReaction.value = myReactionData.reaction_type;
    }


    return { like, dislike, myReaction } as const;
}