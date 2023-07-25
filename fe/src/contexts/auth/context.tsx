import { createContext, useContext } from "react";

type AuthContext = {
  user: {
    username: string;
  };
  token: string | null;
  login: (username: string, password: string) => Promise<void>;
  logout: () => void;
};

const context = createContext<AuthContext>({
  user: {
    username: "",
  },
  token: null,
  login: async () => {
    return undefined;
  },
  logout: () => {},
});

export default context;

export const useAuth = () => {
  return useContext(context);
};
