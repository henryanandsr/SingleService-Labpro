import { createContext, useContext } from "react";

type LoginResponseData = {
  user: {
    username: string;
    name: string;
  };
  token: string;
};

type AuthContext = {
  user: {
    username: string;
  };
  token: string | null;
  login: (username: string, password: string) => Promise<LoginResponseData>;
  logout: () => void;
};

const context = createContext<AuthContext>({
  user: {
    username: "",
  },
  token: null,
  login: async () => {
    return {
      user: {
        username: "",
        name: "",
      },
      token: "",
    };
  },
  logout: () => {},
});

export default context;

export const useAuth = () => {
  return useContext(context);
};
