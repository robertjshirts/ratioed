import { betterAuth } from "npm:better-auth";
import { drizzleAdapter } from "npm:better-auth/adapters/drizzle";
import { db } from "./db/index.ts";

export const auth = betterAuth({
  database: drizzleAdapter(db, { provider: "pg" }),
  emailAndPassword: {
    enabled: true,
  },
});
