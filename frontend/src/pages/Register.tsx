import { useForm } from "react-hook-form";
import { RegisterSchema, type RegisterFormData } from "../schemas/register";
import { zodResolver } from "@hookform/resolvers/zod";
import { HiEyeOff, HiEye } from "react-icons/hi";
import { Link } from "react-router";
import FieldError from "../components/FieldError";
import { useState } from "react";
import { ProjoPath } from "../constants/path";

const Register = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormData>({ resolver: zodResolver(RegisterSchema) });

  const [isHidePassword, setIsHidePassword] = useState<boolean>(true);
  const [isHideConfirmPassword, setIsHideConfirmPassword] =
    useState<boolean>(true);

  const onSubmit = (data: RegisterFormData) => {
    console.log("Valid form data:", data);
  };

  return (
    <main className="h-dvh w-full flex justify-center items-center">
      <div className="card bg-base-100 sm:w-1/2 md:w-96 shadow-md">
        <form className="card-body" onSubmit={handleSubmit(onSubmit)}>
          <h2 className="card-title">Register your account</h2>
          <div className="m-2 flex flex-col justify-center items-center gap-3">
            <div className="w-full flex flex-col gap-1">
              <label htmlFor="firstName" className="self-start">
                First name
              </label>
              <input
                id="firstName"
                type="text"
                placeholder="First name"
                className={`input w-full ${errors.firstName ? "input-error" : "input-bordered"}`}
                {...register("firstName")}
              />
              <FieldError message={errors.firstName?.message} />
            </div>
            <div className="w-full flex flex-col gap-1">
              <label htmlFor="lastName" className="self-start">
                Last name
              </label>
              <input
                id="lastName"
                type="text"
                placeholder="Last name"
                className={`input w-full ${errors.lastName ? "input-error" : "input-bordered"}`}
                {...register("lastName")}
              />
              <FieldError message={errors.lastName?.message} />
            </div>
            <div className="w-full flex flex-col gap-1">
              <label htmlFor="username" className="self-start">
                Username
              </label>
              <input
                autoComplete="username"
                id="username"
                type="text"
                placeholder="Username"
                className={`input w-full ${errors.username ? "input-error" : "input-bordered"}`}
                {...register("username")}
              />
              <FieldError message={errors.username?.message} />
            </div>
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
              <label htmlFor="password" className="self-start">
                Password
              </label>
              <div className="relative">
                <input
                  autoComplete="current-password new-password"
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

            <div className="w-full flex flex-col gap-1">
              <label htmlFor="confirmPassword" className="self-start">
                Confirm Password
              </label>
              <div className="relative">
                <input
                  autoComplete="new-password"
                  id="confirmPassword"
                  type={isHideConfirmPassword ? "password" : "text"}
                  placeholder="Confirm Password"
                  className={`input w-full pr-10 ${errors.confirmPassword ? "input-error" : "input-bordered"}`}
                  {...register("confirmPassword")}
                />
                <label className="swap swap-rotate absolute right-3 top-1/2 -translate-y-1/2 cursor-pointer">
                  <input
                    type="checkbox"
                    checked={isHideConfirmPassword}
                    onChange={() =>
                      setIsHideConfirmPassword(!isHideConfirmPassword)
                    }
                    className="sr-only"
                  />
                  <HiEyeOff className="swap-on text-md text-gray-400" />
                  <HiEye className="swap-off text-md text-gray-400" />
                </label>
              </div>
              <FieldError message={errors.confirmPassword?.message} />
            </div>
          </div>
          <div className="card-actions justify-center">
            <button type="submit" className="w-full btn btn-primary">
              Register
            </button>
          </div>
          <p className="mt-2 text-center">
            Already have an account?{" "}
            <Link to={ProjoPath.LOGIN} className="text-blue-400">
              Login
            </Link>
          </p>
        </form>
      </div>
    </main>
  );
};

export default Register;
