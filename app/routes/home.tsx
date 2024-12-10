import { db } from "~/db/db";
import { playingWithNeon } from "~/db/migrations/schema";
import type { Route } from "./+types/home";
import { Form } from "react-router";

export async function loader() {
  const entries = await db.select().from(playingWithNeon);
  return entries;
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const body = formData.get("body");
  if (!body) return;

  await db.insert(playingWithNeon).values({ body: body });
}

export default function Home({ loaderData: posts }: Route.ComponentProps) {
  return (
    <div className="w-dvw h-dvh flex justify-center items-center">
      <div className="p-6 border-[1px] rounded-xl border-neutral-700">
        <h2 className="text-xl font-bold">Posts</h2>
        <ul>
          {posts.map((post) => (
            <li key={post.id}>{post.body}</li>
          ))}
        </ul>
        <Form method="post">
          <input
            type="text"
            name="body"
            className="bg-neutral-900 border-neutral-700 border rounded-md px-2"
          />
          <button
            type="submit"
            className="ms-4 px-2 rounded-md bg-white text-black"
          >
            Submit
          </button>
        </Form>
      </div>
    </div>
  );
}
