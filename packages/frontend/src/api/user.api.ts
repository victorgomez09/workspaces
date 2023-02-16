import axios from "axios";
import { API_PARAMS } from "../constants/api.contant";

export const getMe = () => {
  return axios.get(`${API_PARAMS.url}/users/me`, {
    withCredentials: true,
  });
};
