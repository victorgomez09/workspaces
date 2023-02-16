import { useQuery } from "react-query";
import { User } from "../models";

export const useMe = (): User | null => {
  const { data, isSuccess } = useQuery<User>("getMe");

  if (isSuccess) return data;

  return null;
};
