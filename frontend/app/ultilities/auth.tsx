import { Outlet, redirect, useLocation, Navigate } from "react-router";
import type { LoaderFunctionArgs } from "react-router";

export async function loader({ request }: LoaderFunctionArgs) {
  if (typeof window === "undefined") {
    return null;
  }
  const token = localStorage.getItem("token");
  const location = useLocation();

  if (!token) {
    return <Navigate to="/" state={{ from: location }} replace />;
  }

  return null;
}

export default function RequireAuth() {
  return <Outlet />;
}
