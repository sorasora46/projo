import { useLocation } from "react-router";
import Topbar from "../components/Topbar";
import Sidebar from "../components/Sidebar";
import { useEffect, useRef, useState } from "react";
import { NewProjectSchema, type NewProjectFormData } from "../schemas/new-project";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import FieldError from "../components/FieldError";
import { api } from "../apis/api";
import type { Project } from "../types/project";

const Home = () => {
  const { pathname } = useLocation();

  const {
    handleSubmit,
    formState: { errors },
    register,
  } = useForm<NewProjectFormData>({ resolver: zodResolver(NewProjectSchema) });


  const modalRef = useRef<HTMLDialogElement>(null);

  const handleNewProjectButtonClick = () => {
    const modal = modalRef.current;
    if (modal) {
      modal.showModal();
    }
  }

  const handleCloseNewProjectModal = () => {
    const modal = modalRef.current;
    if (modal) {
      modal.close();
    }
  }

  const onSubmit = async (data: NewProjectFormData) => {
    try {
      const response = await api.post("/project/", data);
      console.log("submitted data:", data);
      console.log(response);
      setProjects([...projects, { ...data }]);
    } catch (error) {
      console.error(error);
    } finally {
      const modal = modalRef.current;
      if (modal) {
        modal.close();
      }
    }
  }

  const [projects, setProjects] = useState<Project[]>([]);
  useEffect(() => {
    const fetchProjects = async () => {
      // Simulate an API delay
      const response = await api.get("/project/");
      const projects = response.data.result;
      console.log("response data:", projects);
      setProjects(projects);
    }

    fetchProjects();
  }, []);

  return (
    <div className="h-dvh w-full flex flex-col sm:flex-row bg-base-200">
      <Sidebar pathname={pathname} />
      <div className="flex flex-col w-full">
        <Topbar />
        <main className="flex-1 p-3">
          <button className="btn btn-primary" onClick={handleNewProjectButtonClick}>New project</button>
          <dialog ref={modalRef} id="new-project-modal" className="modal">
            <div className="modal-box">
              <h3 className="font-bold text-lg my-2">Create a new project</h3>
              <form className="self-start flex flex-col gap-5" onSubmit={handleSubmit(onSubmit)}>
                <div>
                  <label htmlFor="projectName">
                    Project Name
                  </label>
                  <input
                    id="projectName"
                    type="text"
                    placeholder="Unnamed Project"
                    className={`input w-full ${errors.name ? "input-error" : "input-bordered"}`}
                    {...register("name")}
                  />
                  <FieldError message={errors.name?.message} />
                </div>
                <div>
                  <label htmlFor="projectDescription">
                    Project's description
                  </label>
                  <input
                    id="projectDescription"
                    type="text"
                    placeholder="Description"
                    className={"input w-full input-bordered"}
                    {...register("description")}
                  />
                  <FieldError message={errors.description?.message} />
                </div>
                <div className="self-end flex gap-2">
                  <button type="button" className="btn btn-outline btn-secondary" onClick={handleCloseNewProjectModal}>Cancel</button>
                  <button type="submit" className="btn btn-primary">Create</button>
                </div>
              </form>
            </div>
          </dialog>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 mt-6">
            {projects.map((project, idx) => {
              const tasks = project.projectTasks;
              let taskRatioString = "No task assigned";
              if (tasks) {
                const totalTasks = tasks.length;
                if (totalTasks > 0) {
                  const completedTasks = tasks.filter((task) => task.status === 'completed').length;
                  taskRatioString = `${completedTasks} / ${totalTasks} completed`
                }
              }

              return (
                <div key={project.id! + idx} className="card bg-base-100 shadow-md p-5 border border-base-300">
                  <h2 className="text-xl font-semibold mb-2">{project.name}</h2>
                  <p className="text-sm text-gray-600 mb-3">{project.description}</p>
                  <div className="text-sm mb-2">
                    <strong>Tasks:</strong> {taskRatioString}
                  </div>
                  <div className="flex flex-wrap gap-2">
                    {tasks && tasks.map((task) => (
                      <span
                        key={task.id}
                        className={`badge ${task.status === 'completed'
                          ? 'badge-success'
                          : task.status === 'in-progress'
                            ? 'badge-warning'
                            : 'badge-neutral'
                          }`}
                      >
                        {task.name}
                      </span>
                    ))}
                  </div>
                </div>
              );
            })}
          </div>
        </main>
      </div>
    </div>
  );
};

export default Home;
