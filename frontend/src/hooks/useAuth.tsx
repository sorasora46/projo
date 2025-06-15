import {
  createContext,
  useContext,
  useEffect,
  useState,
  type FC,
  type ReactNode,
} from "react";

interface IAuthContext {
  authenticated: boolean;
  loading: boolean;
}

const defaultAuthContext: IAuthContext = {
  authenticated: false,
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
    // Simulated API call - replace with real one
    const fetchAuthStatus = async () => {
      try {
        // Example: const response = await fetch("/api/auth/status");
        // const data = await response.json();
        const data = await new Promise<{ authenticated: boolean }>((resolve) =>
          setTimeout(() => resolve({ authenticated: true }), 1000),
        );

        setAuthenticated(data.authenticated);
      } catch (error) {
        console.error("Failed to fetch auth status", error);
      } finally {
        setLoading(false);
      }
    };

    fetchAuthStatus();
  }, []);
  return (
    <AuthContext.Provider value={{ authenticated, loading }}>
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
