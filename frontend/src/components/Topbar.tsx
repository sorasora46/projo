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

const Topbar = () => {
  const [isCalendarHover, setIsCalendarHover] = useState(false);
  const [isBellHover, setIsBellHover] = useState(false);
  const [isProfileOpen, setIsProfileOpen] = useState(false);
  return (
    <header className="w-full flex justify-between items-center px-3">
      <h1>Hello, firstName lastName</h1>
      <div className="flex justify-center items-center gap-2">
        <ul className="flex justify-center items-center gap-2 text-lg">
          <li>
            <Link
              to=""
              className="tooltip tooltip-bottom tooltip-primary inline-flex items-center justify-center w-10 h-10"
              data-tip="Calendar"
              onMouseEnter={() => setIsCalendarHover(true)}
              onMouseLeave={() => setIsCalendarHover(false)}
            >
              {isCalendarHover ? <HiCalendarDays /> : <HiOutlineCalendarDays />}
            </Link>
          </li>
          <li>
            <Link
              to=""
              className="tooltip tooltip-bottom tooltip-primary inline-flex items-center justify-center w-10 h-10"
              data-tip="Notification"
              onMouseEnter={() => setIsBellHover(true)}
              onMouseLeave={() => setIsBellHover(false)}
            >
              {isBellHover ? <HiBell /> : <HiOutlineBell />}
            </Link>
          </li>
        </ul>
        <button
          className="btn btn-link text-xl text-black"
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
  );
};

export default Topbar;
