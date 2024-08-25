import type { Database } from "~/types/database";

export const useProfileStore = defineStore("profile", () => {
  const id = useCookie("id");
  const username = useCookie("username");
  const email = useCookie("email");
  const avatarUrl = useCookie("avatarUrl");
  const role = useCookie("role");

  async function signIn() {
    clearProfile();
    const supabase = useSupabaseClient<Database>();
    const user = useSupabaseUser();

    const { data, error } = await supabase
      .from("profiles")
      .select(`id, username, avatar_url, role`)
      .eq("id", user.value.id)
      .single();

    if (error) {
      return;
    }

    id.value = data.id;
    username.value = data.username;
    email.value = user.value.email || "";
    avatarUrl.value = data.avatar_url || "";
    role.value = data.role;
  }

  async function signOut() {
    const supabase = useSupabaseClient<Database>();
    await supabase.auth.signOut();
    clearProfile();
    navigateTo("/login");
  }

  function clearProfile() {
    id.value = "";
    username.value = "";
    email.value = "";
    avatarUrl.value = "";
  }

  return {
    id,
    username,
    email,
    avatarUrl,
    role,
    signIn,
    signOut,
    clearProfile,
  };
});
