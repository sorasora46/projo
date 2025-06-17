import { useState, useEffect, type FC } from "react";
import { AiFillProject, AiOutlineProject } from "react-icons/ai";
import { MdDelete, MdDeleteOutline } from "react-icons/md";
import { RiInboxArchiveFill, RiInboxArchiveLine } from "react-icons/ri";
import { ProjoPath } from "../constants/path";
import { Link } from "react-router";
import { FiMenu } from "react-icons/fi"; // Hamburger icon

interface SidebarProps {
  pathname: string;
}

const Sidebar: FC<SidebarProps> = ({ pathname }) => {
  const [isProjectHover, setIsProjectHover] = useState(false);
  const [isArchivedHover, setIsArchivedHover] = useState(false);
  const [isDeletedHover, setIsDeletedHover] = useState(false);
  const [isMobile, setIsMobile] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  // Track screen size
  useEffect(() => {
    const handleResize = () => {
      setIsMobile(window.innerWidth <= 1024); // iPad width and below
    };

    handleResize(); // Initial check
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  const renderNavLinks = () => (
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
  );

  return (
    <aside className="relative">
      {/* Hamburger for mobile */}
      {isMobile && (
        <button
          className="p-3 text-2xl fixed top-4 left-4 z-50 bg-white rounded shadow-md"
          onClick={() => setIsMenuOpen(!isMenuOpen)}
        >
          <FiMenu />
        </button>
      )}

      {/* Sidebar */}
      <div
        className={`fixed top-0 left-0 h-full bg-white shadow-md transition-transform duration-300 z-40 ${
          isMobile
            ? isMenuOpen
              ? "translate-x-0"
              : "-translate-x-full"
            : "translate-x-0"
        }`}
        style={{ width: isMobile ? "200px" : "auto" }}
      >
        <div className="card-body p-4">
          <h2 className="card-title text-2xl mb-4">Projo</h2>
          <nav>{renderNavLinks()}</nav>
        </div>
      </div>

      {/* Overlay when sidebar is open on mobile */}
      {isMobile && isMenuOpen && (
        <div
          className="fixed inset-0 bg-black opacity-50 z-30"
          onClick={() => setIsMenuOpen(false)}
        />
      )}
    </aside>
  );
};

export default Sidebar;
