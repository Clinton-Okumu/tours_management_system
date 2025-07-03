"use client";
import logo from "@/assets/logo.png";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React, { useEffect, useState } from "react";

const NavbarLinks = [
  { id: 1, title: "Home", link: "/" },
  { id: 2, title: "Tours", link: "/tours" },
  { id: 3, title: "Features", link: "/features" },
  { id: 4, title: "About", link: "/about" },
  { id: 5, title: "Contact", link: "/contact" },
];

const Navbar = () => {
  const [isMobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);
  const pathname = usePathname();

  // Handle scroll effect
  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 20) {
        setScrolled(true);
      } else {
        setScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // Close mobile menu when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      const target = event.target as HTMLElement;
      if (isMobileMenuOpen && !target.closest(".nav-container")) {
        setMobileMenuOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [isMobileMenuOpen]);

  // Check if link is active
  const isActiveLink = (link: string) => {
    if (link === "/" && pathname === "/") return true;
    if (link !== "/" && pathname.startsWith(link)) return true;
    return false;
  };

  return (
    <div
      className={`w-full mx-auto px-4 py-4 flex justify-between items-center sticky top-0 z-50 transition-all duration-300 nav-container ${
        scrolled
          ? "bg-white shadow-md"
          : "bg-white/95 shadow-[0_1px_1px_rgba(0,0,0,0.05)]"
      }`}
    >
      {/* Logo section */}
      <div className="flex items-center gap-3">
        <Image
          src={logo}
          alt="FitJourney logo"
          width={50}
          height={50}
          className="w-[50px]"
          priority
        />
        <p className="font-bold text-xl text-gray-800">
          Tour<span className="text-orange-500">Vista</span>
        </p>
      </div>

      {/* Desktop Menu */}
      <div className="hidden md:block">
        <ul className="flex gap-6">
          {NavbarLinks.map((navLink) => (
            <li key={navLink.id}>
              <Link
                className={`font-medium transition-colors duration-200 relative group ${
                  isActiveLink(navLink.link)
                    ? "text-orange-500"
                    : "text-gray-800 hover:text-orange-500"
                }`}
                href={navLink.link}
              >
                {navLink.title}
                <span
                  className={`absolute bottom-0 left-0 h-0.5 bg-orange-400 transition-all duration-300 ${
                    isActiveLink(navLink.link)
                      ? "w-full"
                      : "w-0 group-hover:w-full"
                  }`}
                ></span>
              </Link>
            </li>
          ))}
        </ul>
      </div>

      {/* Button section */}
      <div className="hidden md:flex space-x-4">
        {["Login", "Signup"].map((label) => (
          <button
            key={label}
            className="bg-orange-500 hover:bg-orange-600 px-6 py-2 rounded-lg text-white font-medium transition-all duration-300 shadow-sm hover:shadow-md"
          >
            {label}
          </button>
        ))}
      </div>

      {/* Mobile Menu Button */}
      <div className="md:hidden flex items-center">
        <button
          className="p-2 text-gray-700 hover:text-orange-500 transition-colors focus:outline-none"
          onClick={() => setMobileMenuOpen(!isMobileMenuOpen)}
          aria-label={isMobileMenuOpen ? "Close menu" : "Open menu"}
        >
          {isMobileMenuOpen ? (
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          ) : (
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          )}
        </button>
      </div>

      {/* Mobile Menu */}
      {isMobileMenuOpen && (
        <div className="absolute top-full left-0 right-0 bg-white p-5 shadow-lg md:hidden border-t border-gray-100 animate-fade-in">
          <ul className="flex flex-col gap-4">
            {NavbarLinks.map((navLink) => (
              <li key={navLink.id} className="border-b border-gray-100 pb-2">
                <Link
                  className={`block transition-colors text-lg font-medium ${
                    isActiveLink(navLink.link)
                      ? "text-orange-500"
                      : "text-gray-800 hover:text-orange-500"
                  }`}
                  href={navLink.link}
                  onClick={() => setMobileMenuOpen(false)}
                >
                  {navLink.title}
                </Link>
              </li>
            ))}
            <li className="pt-2 space-y-2">
              <button className="w-full bg-orange-400 hover:bg-orange-500 active:scale-95 px-6 py-3 rounded-lg text-white font-semibold text-center shadow-md hover:shadow-lg transition-all duration-300">
                Login
              </button>
              <button className="w-full bg-white hover:bg-orange-100 border border-orange-300 px-6 py-3 rounded-lg text-orange-500 font-semibold text-center shadow-sm hover:shadow-md transition-all duration-300">
                Signup
              </button>
            </li>
          </ul>
        </div>
      )}
    </div>
  );
};

export default Navbar;
