import { useLocation } from "react-router";
import Topbar from "../components/Topbar";
import Sidebar from "../components/Sidebar";

const Home = () => {
  const { pathname } = useLocation();

  return (
    <div className="h-dvh w-full flex flex-col md:flex-row bg-base-200">
      <Sidebar pathname={pathname} />
      <div className="flex flex-col w-full">
        <Topbar />
        <main className="flex-1 p-3">
          <p>Main content area</p>
          <p>{pathname}</p>
        </main>
      </div>
    </div>
  );
};

export default Home;
