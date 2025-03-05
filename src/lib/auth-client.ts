import { createAuthClient } from "npm:better-auth/react";

export const authClient = createAuthClient({
  baseURL: "http://localhost:3000",
});
