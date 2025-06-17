import { useLocation, Link } from "react-router";
import Topbar from "../components/Topbar";
import { AiFillProject, AiOutlineProject } from "react-icons/ai";
import { useState } from "react";
import { ProjoPath } from "../constants/path";
import { MdDelete, MdDeleteOutline } from "react-icons/md";
import { RiInboxArchiveFill, RiInboxArchiveLine } from "react-icons/ri";

const Home = () => {
  const [isProjectHover, setIsProjectHover] = useState<boolean>(false);
  const [isArchivedHover, setIsArchivedHover] = useState<boolean>(false);
  const [isDeletedHover, setIsDeletedHover] = useState<boolean>(false);
  const { pathname } = useLocation();

  return (
    <div className="h-dvh w-full flex">
      <aside className="h-full bg-blue-500 p-4">
        <h2 className="text-2xl">Projo</h2>
        <nav>
          <ul className="menu text-xl">
            <li>
              <Link
                to={ProjoPath.HOME}
                className="tooltip tooltip-right tooltip-primary"
                data-tip="Home"
                onMouseEnter={() => setIsProjectHover(true)}
                onMouseLeave={() => setIsProjectHover(false)}
              >
                {isProjectHover || pathname === ProjoPath.HOME ? (
                  <AiFillProject />
                ) : (
                  <AiOutlineProject />
                )}
              </Link>
              <Link
                to={ProjoPath.ARCHIVED}
                className="tooltip tooltip-right tooltip-primary"
                data-tip="Archived"
                onMouseEnter={() => setIsArchivedHover(true)}
                onMouseLeave={() => setIsArchivedHover(false)}
              >
                {isArchivedHover || pathname === ProjoPath.ARCHIVED ? (
                  <RiInboxArchiveFill />
                ) : (
                  <RiInboxArchiveLine />
                )}
              </Link>
              <Link
                to={ProjoPath.DELETED}
                className="tooltip tooltip-right tooltip-primary"
                data-tip="Deleted"
                onMouseEnter={() => setIsDeletedHover(true)}
                onMouseLeave={() => setIsDeletedHover(false)}
              >
                {isDeletedHover || pathname === ProjoPath.DELETED ? (
                  <MdDelete />
                ) : (
                  <MdDeleteOutline />
                )}
              </Link>
            </li>
          </ul>
        </nav>
      </aside>
      <div className="w-full flex flex-col">
        <Topbar />
        <main className="flex-1 bg-gray-100 p-4">
          <p>Main content area</p>
          <p>{pathname}</p>
        </main>
      </div>
    </div>
  );
};

export default Home;
