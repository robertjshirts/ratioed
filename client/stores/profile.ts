import type { Database } from "~/types/database";

export const useProfileStore = defineStore("profile", () => {
  const id = useCookie("id");
  const username = useCookie("username");
  const email = useCookie("email");
  const avatarUrl = useCookie("avatarUrl");

  async function signIn() {
    clearProfile();
    const supabase = useSupabaseClient<Database>();
    const user = useSupabaseUser();

    const { data, error } = await supabase
      .from("profiles")
      .select(`id, username, avatar_url`)
      .eq("id", user.value.id)
      .single();

    console.log(data, error);

    if (data) {
      id.value = data.id;
      username.value = data.username;
      email.value = user.value.email || "";
      avatarUrl.value = data.avatar_url || "";
    }
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

  return { id, username, email, avatarUrl, signIn, signOut, clearProfile };
});
