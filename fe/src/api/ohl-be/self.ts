import { Response, User } from "@/types/models";
import {api} from ".";

type SelfData = User

const selfAPI = async ()  => {
  const response = await api.get<Response<SelfData>>("/self");
  return response.data;
}

export default selfAPI;