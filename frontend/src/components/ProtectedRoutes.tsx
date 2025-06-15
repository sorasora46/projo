import { Navigate, Outlet } from "react-router";
import { useAuth } from "../hooks/useAuth";

const ProtectedRoute = () => {
  const { authenticated, loading } = useAuth();

  if (loading) {
    return <p>loading . . .</p>;
  }

  return authenticated ? <Outlet /> : <Navigate to="/login" replace />;
};

export default ProtectedRoute;
