import type { Route } from "./+types/home";
import { redirect } from "react-router";
import { Welcome } from "../welcome/welcome";

export async function loader() {
  if (typeof window === "undefined") {
    return null;
  }

  const token = localStorage.getItem("token");

  if (!token) {
    console.log("loader JWT token:", token);
    throw redirect("/dashboard");
  }
  return null;
}

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
