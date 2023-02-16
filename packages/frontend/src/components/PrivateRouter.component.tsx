import { Navigate, Outlet } from "react-router-dom";
import { useMe } from "../hooks";

export function ProtectedRoute() {
  const user = useMe();

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
}
