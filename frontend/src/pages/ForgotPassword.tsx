import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import FieldError from "../components/FieldError";
import {
  ForgotPasswordSchema,
  type ForgotPasswordFormData,
} from "../schemas/forgot-password";
import { Link } from "react-router";
import { ProjoPath } from "../constants/path";

const ForgotPassword = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<ForgotPasswordFormData>({
    resolver: zodResolver(ForgotPasswordSchema),
  });

  const onSubmit = (data: ForgotPasswordFormData) => {
    // TODO: implement forgot password endpoint
    console.log("Valid form data:", data);
  };

  return (
    <main className="h-dvh w-full flex justify-center items-center">
      <div className="card bg-base-100 sm:w-1/2 md:w-96 shadow-md">
        <form className="card-body" onSubmit={handleSubmit(onSubmit)}>
          <h2 className="card-title">Change your password</h2>
          <p>Enter your email to reset the password</p>
          <div className="m-2 flex flex-col justify-center items-center gap-3">
            <div className="w-full flex flex-col gap-1">
              <label htmlFor="email" className="self-start">
                Email
              </label>
              <input
                id="email"
                type="text"
                placeholder="example@mail.com"
                className={`input w-full ${errors.email ? "input-error" : "input-bordered"}`}
                {...register("email")}
              />
              <FieldError message={errors.email?.message} />
            </div>
          </div>
          <div className="card-actions justify-center">
            <button type="submit" className="w-full btn btn-primary">
              Send
            </button>
          </div>
          <p className="text-center mt-3">
            Remember now?{" "}
            <Link to={ProjoPath.LOGIN} className="text-blue-400">
              Login
            </Link>
          </p>
        </form>
      </div>
    </main>
  );
};

export default ForgotPassword;
