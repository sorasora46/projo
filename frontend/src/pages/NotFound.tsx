import { useNavigate } from "react-router";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <main className="p-3 h-dvh w-full flex flex-col justify-center items-center text-center gap-5">
      <p className="text-8xl text-blue-400">404</p>
      <div className="flex flex-col justify-center items-center text-xl gap-3">
        <h1 className="text-3xl">Page not found</h1>
        <p>Sorry, We can't find the page you're looking for.</p>
      </div>
      <button
        className="btn btn-outline btn-info hover:text-white"
        onClick={() => navigate("/")}
      >
        Go Home
      </button>
    </main>
  );
};

export default NotFound;
