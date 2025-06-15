import z from "zod/v4";

export type RegisterFormData = z.infer<typeof RegisterSchema>;

export const RegisterSchema = z
  .object({
    firstName: z.string().min(2, "First name must be at least 2 characters"),
    lastName: z.string().min(2, "Last name must be at least 2 characters"),
    username: z.string().min(6, "Username must be at least 6 characters"),
    email: z.email("Invalid/Empty email address"),
    password: z.string().min(6, "Password must be at least 6 characters"),
    confirmPassword: z
      .string()
      .min(6, "Password must be at least 6 characters"),
  })
  .refine((schema) => schema.password === schema.confirmPassword, {
    message: "Password must match",
    path: ["confirmPassword"],
  });
