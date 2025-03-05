CREATE TYPE "public"."reaction" AS ENUM('like', 'dislike');--> statement-breakpoint
CREATE TABLE "posts" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"authorId" uuid NOT NULL,
	"parentId" uuid,
	"body" varchar(400),
	"updated_at" timestamp,
	"created_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "reactions" (
	"postId" uuid NOT NULL,
	"userId" uuid NOT NULL,
	"reaction" "reaction" NOT NULL,
	"updated_at" timestamp,
	"created_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "users" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"username" varchar(20) NOT NULL,
	"bio" varchar(200),
	"updated_at" timestamp,
	"created_at" timestamp DEFAULT now(),
	"deleted_at" timestamp,
	CONSTRAINT "username" UNIQUE("username")
);
--> statement-breakpoint
ALTER TABLE "posts" ADD CONSTRAINT "posts_authorId_fk" FOREIGN KEY ("authorId") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "posts" ADD CONSTRAINT "post_parentId_fk" FOREIGN KEY ("parentId") REFERENCES "public"."posts"("id") ON DELETE no action ON UPDATE no action;