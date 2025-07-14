import { z } from "zod/v4";

export type LoginFormData = z.infer<typeof LoginSchema>;

export const LoginSchema = z.object({
  // email: z.email("Invalid/Empty email address"),
  username: z.string().min(6, "Username must be at least 6 characters"),
  password: z.string().min(6, "Password must be at least 6 characters"),
});
