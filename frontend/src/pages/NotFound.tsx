import { useNavigate } from "react-router";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <main className="hero bg-base-200 min-h-screen">
      <div className="hero-content flex-col lg:flex-row-reverse text-center">
        <div className="max-w-md">
          <h1 className="text-6xl font-extrabold text-error mb-4">404</h1>
          <h2 className="text-2xl font-semibold mb-2">Oops! Page not found</h2>
          <p className="py-4">
            The page you’re looking for doesn’t exist or has been moved. Let’s
            get you back on track!
          </p>
          <div className="flex justify-center gap-4 mt-6">
            <button className="btn btn-primary" onClick={() => navigate("/")}>
              🏠 Go Home
            </button>
            <button
              className="btn btn-outline"
              onClick={() => navigate("/contact")}
            >
              📩 Contact Support
            </button>
          </div>
        </div>
      </div>
    </main>
  );
};

export default NotFound;
