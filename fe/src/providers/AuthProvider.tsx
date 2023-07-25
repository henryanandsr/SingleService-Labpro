import { LoadingSpinner } from "@/components";
import { AuthContext } from "@/contexts";
import { useAPI } from "@/contexts";
import { User } from "@/types/models";
import { AxiosError } from "axios";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

interface AuthProviderProps {
  children: React.ReactNode;
}

export default function AuthProvider({ children }: AuthProviderProps) {
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [user, setUser] = useState<User>({
    username: "",
  });
  const location = useLocation();
  const navigate = useNavigate();
  const { api, token, setToken } = useAPI();

  const login = async (username: string, password: string) => {
    const res = await api.login(username, password);
    if (res.data) {
      setToken(res.data.token);
      navigate("/barang");
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
    setUser({
      username: "",
    });
    navigate("/login");
  };

  useEffect(() => {
    if (location.pathname === "/config") {
      setIsLoading(false);
      return;
    }
    // token check
    if (token) {
      localStorage.setItem("token", token);
    } else {
      setToken(localStorage.getItem("token") || null);
    }

    // don't check token on login page
    if (location.pathname === "/login") {
      // if token exists, redirect to home
      if (token) {
        navigate("/barang");
        setIsLoading(false);
        return;
      }
      setIsLoading(false);
    } else {
      // if token doesn't exist, redirect to login
      if (!localStorage.getItem("token")) {
        navigate("/login");
        return;
      }
      // send token to backend to get user data
      api
        .self()
        .then((res) => {
          if (!res.data) {
            throw new AxiosError("Unauthorized", "401");
          }
          // set user data
          setUser(res.data);
          setIsLoading(false);
        })
        .catch((err) => {
          if (err.code === "401") {
            localStorage.removeItem("token");
            setToken(null);
            navigate("/login");
          }
        });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [token, location.pathname]);

  if (isLoading) {
    return <LoadingSpinner />;
  }

  return (
    <AuthContext.Provider
      value={{
        user: user,
        token: null,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
