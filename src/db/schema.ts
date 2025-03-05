import {
  foreignKey,
  pgEnum,
  pgTable,
  timestamp,
  uuid,
  varchar,
} from "drizzle-orm/pg-core";

const timestamps = {
  updated_at: timestamp(),
  created_at: timestamp().defaultNow(),
  deleted_at: timestamp(),
};

export const usersTable = pgTable("users", {
  id: uuid().defaultRandom().primaryKey(),
  username: varchar({ length: 20 }).notNull().unique("username"), // More than twitter
  bio: varchar({ length: 200 }), // More than twitter
  ...timestamps,
});

export const postsTable = pgTable(
  "posts",
  {
    id: uuid().defaultRandom().primaryKey(),
    authorId: uuid().notNull(),
    parentId: uuid(),
    body: varchar({ length: 400 }), // More than twitter
    ...timestamps,
  },
  (table) => [
    foreignKey({
      // posts -> users
      columns: [table.authorId],
      foreignColumns: [usersTable.id],
      name: "posts_authorId_fk",
    }),
    foreignKey({
      // posts -> posts (for comments)
      columns: [table.parentId],
      foreignColumns: [table.id],
      name: "post_parentId_fk",
    }),
  ],
);

export const reactionEnum = pgEnum("reaction", ["like", "dislike"]);

export const reactionsTable = pgTable("reactions", {
  postId: uuid().notNull(),
  userId: uuid().notNull(),
  reaction: reactionEnum().notNull(),
  ...timestamps,
});
