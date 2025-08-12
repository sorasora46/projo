import { z } from "zod/v4";

export type NewProjectFormData = z.infer<typeof NewProjectSchema>;

export const NewProjectSchema = z.object({
  name: z
    .string({ error: "Project's name must not be empty" })
    .min(1, "Project's name must not be empty"),
  description: z.optional(z.string()),
});
