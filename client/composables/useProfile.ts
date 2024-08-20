export default function useProfile() {
  const username = useState<string>("username", () => "");
  const avatarUrl = useState<string>("avatarUrl", () => "");

  async function signOut() {
    const supabase = useSupabaseClient();
    const { error } = await supabase.auth.signOut();

    if (error) return;

    username.value = "";
    avatarUrl.value = "";
    reloadNuxtApp();
  }

  return { username, avatarUrl, signOut };
}
