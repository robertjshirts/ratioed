export interface Post {
  id: string;
  username: string;
  pfp: string;
  content: {
    ratioed?: boolean;
    body?: string;
    attachment?: {
      type: string;
      src: string;
    };
  };
  likes: number;
  dislikes: number;
  timestamp: string;
}
