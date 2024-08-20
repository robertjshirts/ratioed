export default function useProfile() {
  return useState<{
    username: string;
    avatarUrl: string | null;
  }>("profile", () => ({
    username: "",
    avatarUrl: "",
  }));
}
