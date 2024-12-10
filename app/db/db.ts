import { drizzle } from "drizzle-orm/neon-http";
import { neon } from "@neondatabase/serverless";

export const db = drizzle({
  client: neon(process.env.DATABASE_URL!),
});
