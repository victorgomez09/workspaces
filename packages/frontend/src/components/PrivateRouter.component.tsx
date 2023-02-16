import { Navigate, Outlet } from "react-router-dom";

export function ProtectedRoute() {
  //   if (!user) {
  //     return <Navigate to="/login" replace />;
  //   }

  return <Outlet />;
}
