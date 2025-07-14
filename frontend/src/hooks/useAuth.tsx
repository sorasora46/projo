import {
  createContext,
  useContext,
  useEffect,
  useState,
  type FC,
  type ReactNode,
} from "react";
import { api } from "../apis/api";

interface IAuthContext {
  authenticated: boolean;
  setAuthenticated: React.Dispatch<React.SetStateAction<boolean>> | null;
  loading: boolean;
}

const defaultAuthContext: IAuthContext = {
  authenticated: false,
  setAuthenticated: null,
  loading: false,
};

const AuthContext = createContext<IAuthContext>(defaultAuthContext);

interface AuthProviderProps {
  children: ReactNode;
}

const AuthProvider: FC<AuthProviderProps> = ({ children }) => {
  const [authenticated, setAuthenticated] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchAuthStatus = async () => {
      try {
        await api.get("/user/check-auth");
        setAuthenticated(true);
      } catch (error) {
        // TODO: show error popup
        setAuthenticated(false);
        console.error("Failed to fetch auth status", error);
      } finally {
        setLoading(false);
      }
    };

    fetchAuthStatus();
  }, []);
  return (
    <AuthContext.Provider value={{ authenticated, setAuthenticated, loading }}>
      {children}
    </AuthContext.Provider>
  );
};

const useAuth = () => {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error("useAuth must be used within AuthProvider");
  }

  return context;
};

export { AuthProvider, useAuth };
