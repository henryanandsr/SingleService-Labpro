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
  return response.data;
}

export default loginAPI;