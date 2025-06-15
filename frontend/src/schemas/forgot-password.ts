import { z } from "zod/v4";

export type ForgotPasswordFormData = z.infer<typeof ForgotPasswordSchema>;

export const ForgotPasswordSchema = z.object({
  email: z.email("Invalid/Empty email address"),
});
