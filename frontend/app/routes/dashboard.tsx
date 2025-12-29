import { redirect } from "react-router";

export async function loader() {
  if (typeof window === "undefined") {
    return null;
  }

  const token = localStorage.getItem("token");

  if (!token) {
    throw redirect("/");
  }

  const res = await fetch("http://localhost:8000/protected", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (res.status === 401 || res.status === 403) {
    localStorage.removeItem("token");
    throw redirect("/");
  }

  return null;
}

export default function Dashboard() {
  return <h1>Dashboard</h1>;
}
