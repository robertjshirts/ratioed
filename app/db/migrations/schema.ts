import { pgTable, serial, text } from "drizzle-orm/pg-core"
import { sql } from "drizzle-orm"



export const playingWithNeon = pgTable("playing_with_neon", {
	id: serial().primaryKey().notNull(),
	body: text().notNull(),
});
