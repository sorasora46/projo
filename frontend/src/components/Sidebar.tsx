import { useState, type FC } from "react";
import { AiFillProject, AiOutlineProject } from "react-icons/ai";
import { ProjoPath } from "../constants/path";
import { MdDelete, MdDeleteOutline } from "react-icons/md";
import { RiInboxArchiveFill, RiInboxArchiveLine } from "react-icons/ri";
import { Link } from "react-router";

interface SidebarProps {
  pathname: string;
}

const Sidebar: FC<SidebarProps> = ({ pathname }) => {
  const [isProjectHover, setIsProjectHover] = useState<boolean>(false);
  const [isArchivedHover, setIsArchivedHover] = useState<boolean>(false);
  const [isDeletedHover, setIsDeletedHover] = useState<boolean>(false);

  return (
    <aside className="h-full">
      <div className="card bg-white shadow-md h-full">
        <div className="card-body p-4">
          <h2 className="card-title text-2xl">Projo</h2>
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
        </div>
      </div>
    </aside>
  );
};

export default Sidebar;
