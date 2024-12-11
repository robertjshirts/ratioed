import { Outlet } from "react-router";
import { Button } from "~/components/ui/button";

export default function Layout() {
  return (
    <div>
      <nav>navbar</nav>
      <Outlet />
      <Button>shadcn button!</Button>
    </div>
  );
}
