import { useState } from "react";
import { HiEye, HiEyeOff } from "react-icons/hi";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { LoginSchema, type LoginFormData } from "../schemas/login";
import { Link, useNavigate } from "react-router";
import FieldError from "../components/FieldError";

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({ resolver: zodResolver(LoginSchema) });

  const [isHidePassword, setIsHidePassword] = useState(true);

  const navigate = useNavigate();

  const onSubmit = (data: LoginFormData) => {
    console.log("Valid form data:", data);
  };

  return (
    <main className="h-dvh w-full flex justify-center items-center">
      <div className="card bg-base-100 sm:w-1/2 md:w-96 shadow-md">
        <form className="card-body" onSubmit={handleSubmit(onSubmit)}>
          <h2 className="card-title">Login to your account</h2>
          <div className="m-2 flex flex-col justify-center items-center gap-3">
            <div className="w-full flex flex-col gap-1">
              <label htmlFor="email" className="self-start">
                Email
              </label>
              <input
                id="email"
                type="text"
                placeholder="Email"
                className={`input w-full ${errors.email ? "input-error" : "input-bordered"}`}
                {...register("email")}
              />
              <FieldError message={errors.email?.message} />
            </div>
            <div className="w-full flex flex-col gap-1">
              <div className="flex justify-between items-center">
                <label htmlFor="password" className="self-start">
                  Password
                </label>
                <Link to="/forgot-password" className="text-xs text-info">
                  Forgot ?
                </Link>
              </div>
              <div className="relative">
                <input
                  id="password"
                  type={isHidePassword ? "password" : "text"}
                  placeholder="Password"
                  className={`input w-full pr-10 ${errors.password ? "input-error" : "input-bordered"}`}
                  {...register("password")}
                />
                <label className="swap swap-rotate absolute right-3 top-1/2 -translate-y-1/2 cursor-pointer">
                  <input
                    type="checkbox"
                    checked={isHidePassword}
                    onChange={() => setIsHidePassword(!isHidePassword)}
                    className="sr-only"
                  />
                  <HiEyeOff className="swap-on text-md text-gray-400" />
                  <HiEye className="swap-off text-md text-gray-400" />
                </label>
              </div>
              <FieldError message={errors.password?.message} />
            </div>
          </div>
          <div className="card-actions justify-center">
            <button type="submit" className="w-full btn btn-primary">
              Login
            </button>
          </div>
          <div className="divider">OR</div>
          <div className="card-actions justify-center">
            <button
              type="button"
              className="w-full btn btn-outline btn-secondary"
              onClick={() => navigate("/register")}
            >
              Sign up
            </button>
          </div>
        </form>
      </div>
    </main>
  );
};

export default Login;
