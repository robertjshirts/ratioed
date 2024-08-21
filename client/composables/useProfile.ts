export default function useProfile() {
  const username = useState<string>("username", () => "");
  const email = useState<string>("email", () => "");
  const avatarUrl = useState<string>("avatarUrl", () => "");

  async function signOut() {
    const supabase = useSupabaseClient();
    const { error } = await supabase.auth.signOut();

    if (error) return;

    username.value = "";
    avatarUrl.value = "";
    reloadNuxtApp();
  }

  return { username, email, avatarUrl, signOut };
}
