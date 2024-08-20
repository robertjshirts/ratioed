export default function useProfile() {
  const username = useState("username", () => "");
  const email = useState("email", () => "");
  const avatarUrl = useState("avatarUrl", () => "");

  return { username, email, avatarUrl };
}
