import { BrowserRouter, Route } from "react-router";
import Home from "./pages/Home.tsx";
import { Routes } from "react-router";
import Login from "./pages/Login.tsx";
import Register from "./pages/Register.tsx";
import NotFound from "./pages/NotFound.tsx";
import { AuthProvider } from "./hooks/useAuth.tsx";
import ProtectedRoutes from "./components/ProtectedRoutes.tsx";
import ForgotPassword from "./pages/ForgotPassword.tsx";
import { ProjoPath } from "./constants/path.ts";

const App = () => {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route element={<ProtectedRoutes />}>
            <Route path={ProjoPath.HOME} element={<Home />} />
            <Route path={ProjoPath.ARCHIVED} element={<Home />} />
            <Route path={ProjoPath.DELETED} element={<Home />} />
          </Route>
          <Route path={ProjoPath.LOGIN} element={<Login />} />
          <Route path={ProjoPath.REGISTER} element={<Register />} />
          <Route
            path={ProjoPath.FORGOT_PASSWORD}
            element={<ForgotPassword />}
          />
          <Route path={ProjoPath.CATCH_ALL} element={<NotFound />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
};

export default App;
