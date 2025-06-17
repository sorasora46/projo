import { useState, type FC } from "react";
import { AiFillProject, AiOutlineProject } from "react-icons/ai";
import { MdDelete, MdDeleteOutline } from "react-icons/md";
import { RiInboxArchiveFill, RiInboxArchiveLine } from "react-icons/ri";
import { ProjoPath } from "../constants/path";
import { Link } from "react-router";
import { FiMenu } from "react-icons/fi";

interface SidebarProps {
  pathname: string;
}

const Sidebar: FC<SidebarProps> = ({ pathname }) => {
  const [isProjectHover, setIsProjectHover] = useState(false);
  const [isArchivedHover, setIsArchivedHover] = useState(false);
  const [isDeletedHover, setIsDeletedHover] = useState(false);

  const renderNavLinks = () => (
    <ul className="menu text-xl">
      <li>
        <Link
          to={ProjoPath.HOME}
          className="md:tooltip md:tooltip-bottom md:tooltip-primary"
          data-tip="Home"
          onMouseEnter={() => setIsProjectHover(true)}
          onMouseLeave={() => setIsProjectHover(false)}
        >
          {isProjectHover || pathname === ProjoPath.HOME ? (
            <AiFillProject />
          ) : (
            <AiOutlineProject />
          )}
          <span className="text-sm md:hidden">Home</span>
        </Link>
      </li>
      <li>
        <Link
          to={ProjoPath.ARCHIVED}
          className="md:tooltip md:tooltip-bottom md:tooltip-primary"
          data-tip="Archived"
          onMouseEnter={() => setIsArchivedHover(true)}
          onMouseLeave={() => setIsArchivedHover(false)}
        >
          {isArchivedHover || pathname === ProjoPath.ARCHIVED ? (
            <RiInboxArchiveFill />
          ) : (
            <RiInboxArchiveLine />
          )}
          <span className="text-sm md:hidden">Archived</span>
        </Link>
      </li>
      <li>
        <Link
          to={ProjoPath.DELETED}
          className="md:tooltip md:tooltip-bottom md:tooltip-primary"
          data-tip="Deleted"
          onMouseEnter={() => setIsDeletedHover(true)}
          onMouseLeave={() => setIsDeletedHover(false)}
        >
          {isDeletedHover || pathname === ProjoPath.DELETED ? (
            <MdDelete />
          ) : (
            <MdDeleteOutline />
          )}
          <span className="text-sm md:hidden">Deleted</span>
        </Link>
      </li>
    </ul>
  );

  return (
    <aside>
      <div className="drawer md:drawer-open">
        <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />
        <div className="drawer-content flex flex-col items-center justify-center">
          <label htmlFor="my-drawer-2" className="btn drawer-button md:hidden">
            <FiMenu />
          </label>
        </div>
        <div className="drawer-side">
          <label
            htmlFor="my-drawer-2"
            aria-label="close sidebar"
            className="drawer-overlay"
          ></label>
          <nav className="menu bg-base-100 h-dvh p-4 flex items-center">
            <label
              htmlFor="my-drawer-2"
              className="btn btn-ghost drawer-button md:hidden"
            >
              Close
            </label>
            {renderNavLinks()}
          </nav>
        </div>
      </div>
    </aside>
  );
};

export default Sidebar;
