import axios from "axios";

export const api = axios.create({
  baseURL: import.meta.env.PROJO_BACKEND_URL,
  withCredentials: true,
});
