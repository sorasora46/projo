import { createContext, useContext, type FC, type ReactNode } from "react";

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
  const authenticated = false;
  const loading = false;
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
