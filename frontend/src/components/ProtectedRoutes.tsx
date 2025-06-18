import { Navigate, Outlet } from "react-router";
import { useAuth } from "../hooks/useAuth";
import { ProjoPath } from "../constants/path";

const ProtectedRoute = () => {
  const { authenticated, loading } = useAuth();

  if (loading) {
    return <p>loading . . .</p>;
  }

  return authenticated ? <Outlet /> : <Navigate to={ProjoPath.LOGIN} replace />;
};

export default ProtectedRoute;
