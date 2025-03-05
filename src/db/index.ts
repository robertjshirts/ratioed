import { neon } from "@neondatabase/serverless";
import { drizzle } from "drizzle-orm/neon-http";

export const sql = neon(Deno.env.get("DATABASE_URL")!);
export const db = drizzle({ client: sql });
