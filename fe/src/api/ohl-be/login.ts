import { Response, User } from "@/types/models";
import { api } from ".";

type LoginData = {
  user: User;
  token: string;
}

const loginAPI = async (username: string, password: string) => {
  const response = await api.post<Response<LoginData>>("/login", {
    username,
    password
  });
  const token = response.data.token;
  if (token) {
    window.localStorage.setItem("token", token);

    api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  }
  return response.data;
}

export default loginAPI;
