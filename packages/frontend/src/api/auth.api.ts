import axios from "axios";
import { API_PARAMS } from "../constants/api.contant";

export type LoginInputs = {
  email: string;
  password: string;
};

export const login = async (data: LoginInputs) => {
  return (
    await axios.post<string>(`${API_PARAMS.url}/auth/login`, data, {
      withCredentials: true,
    })
  ).data;
};
