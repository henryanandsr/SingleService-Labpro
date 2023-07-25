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
  console.log("Login response: ", response);
  const token = response.headers['authorization']?.split(' ')[1];
  if (token) {
    window.localStorage.setItem("token", token);
  }
  return response.data;
}

export default loginAPI;