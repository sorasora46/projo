import { useState } from "react";
import { CgProfile } from "react-icons/cg";
import { HiBell } from "react-icons/hi";
import {
  HiCalendarDays,
  HiOutlineBell,
  HiOutlineCalendarDays,
} from "react-icons/hi2";
import { MdLogout } from "react-icons/md";
import { RiArrowDownSLine, RiArrowUpSLine } from "react-icons/ri";
import { Link } from "react-router";

const Home = () => {
  const [isCalendarHover, setIsCalendarHover] = useState(false);
  const [isBellHover, setIsBellHover] = useState(false);
  const [isProfileOpen, setIsProfileOpen] = useState(false);

  return (
    <div className="h-dvh w-full flex">
      <aside className="h-full bg-blue-500 p-4">
        <h2>sidebar</h2>
        <nav></nav>
      </aside>
      <div className="w-full flex flex-col">
        <header className="w-full flex justify-between items-center bg-blue-500 p-4 pe-3">
          <h1>Hello, firstName lastName</h1>
          <div className="flex items-center gap-2">
            <ul className="flex items-center gap-5 text-lg">
              <li>
                <Link
                  to=""
                  onMouseEnter={() => setIsCalendarHover(true)}
                  onMouseLeave={() => setIsCalendarHover(false)}
                >
                  {isCalendarHover ? (
                    <HiCalendarDays />
                  ) : (
                    <HiOutlineCalendarDays />
                  )}
                </Link>
              </li>
              <li>
                <Link
                  to=""
                  onMouseEnter={() => setIsBellHover(true)}
                  onMouseLeave={() => setIsBellHover(false)}
                >
                  {isBellHover ? <HiBell /> : <HiOutlineBell />}
                </Link>
              </li>
            </ul>
            <button
              className="btn btn-link text-xl"
              popoverTarget="popover-1"
              style={{ anchorName: "--anchor-1" } as React.CSSProperties}
              onClick={() => setIsProfileOpen((prev) => !prev)}
              onBlur={() => setIsProfileOpen(false)}
            >
              <img
                className="mask mask-circle w-8 h-8"
                src="https://img.daisyui.com/images/stock/photo-1567653418876-5bb0e566e1c2.webp"
              />
              {isProfileOpen ? <RiArrowUpSLine /> : <RiArrowDownSLine />}
            </button>
            <ul
              className="dropdown menu bg-base-100 shadow-sm p-0 [&_li>*]:rounded-none [&_li>*]:px-5"
              popover="auto"
              id="popover-1"
              style={{ positionAnchor: "--anchor-1" } as React.CSSProperties}
            >
              <li>
                <Link to="/profile">
                  <CgProfile />
                  Profile
                </Link>
              </li>
              <li>
                <button>
                  <MdLogout />
                  Logout
                </button>
              </li>
            </ul>
          </div>
        </header>
        <main className="flex-1 bg-gray-100 p-4">
          <p>Main content area</p>
        </main>
      </div>
    </div>
  );
};

export default Home;
