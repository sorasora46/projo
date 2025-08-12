export type ProjectTask = {
  id: string;
  name: string;
  status: "pending" | "in-progress" | "completed"; // extend if needed
};

export type Project = {
  id?: string;
  name: string;
  description?: string;
  userId?: string;
  projectTasks?: ProjectTask[];
  createdAt?: string;
  updatedAt?: string;
};
