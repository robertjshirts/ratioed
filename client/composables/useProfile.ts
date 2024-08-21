export default function useProfile() {
  const username = useCookie<string>("username");
  const email = useCookie<string>("email");
  const avatarUrl = useCookie<string>("avatarUrl");

  async function signOut() {
    const supabase = useSupabaseClient();
    const { error } = await supabase.auth.signOut();

    if (error) return;

    username.value = "";
    avatarUrl.value = "";
    navigateTo("/login");
  }

  return { username, email, avatarUrl, signOut };
}
