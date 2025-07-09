"use client";
import logo from "@/assets/logo.png";
import LoginForm from "@/components/auth/Login";
import SignupForm from "@/components/auth/Register";
import Modal from "@/components/auth/Modal";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React, { useEffect, useState } from "react";

const NavbarLinks = [
  { id: 1, title: "Home", link: "/" },
  { id: 2, title: "Tours", link: "/tours" },
  { id: 3, title: "About", link: "/about" },
];

const Navbar = () => {
  const [isMobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);
  const [showLogin, setShowLogin] = useState(false);
  const [showSignup, setShowSignup] = useState(false);
  const [user, setUser] = useState<{ name: string } | null>(null);
  const pathname = usePathname();

  // Scroll detection
  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 20);
    };
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // Click outside to close mobile menu
  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      const target = e.target as HTMLElement;
      if (isMobileMenuOpen && !target.closest(".nav-container")) {
        setMobileMenuOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [isMobileMenuOpen]);

  // Load user from localStorage
  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const isActiveLink = (link: string) =>
    link === pathname || (link !== "/" && pathname.startsWith(link));

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    setUser(null);
  };

  const handleLoginSuccess = () => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
    setShowLogin(false);
  };

  const handleSignupSuccess = () => {
    setShowSignup(false);
  };

  return (
    <>
      {/* Navbar */}
      <div
        className={`w-full mx-auto px-4 py-4 flex justify-between items-center sticky top-0 z-50 transition-all duration-300 nav-container ${
          scrolled
            ? "bg-white shadow-md"
            : "bg-white/95 shadow-[0_1px_1px_rgba(0,0,0,0.05)]"
        }`}
      >
        {/* Logo */}
        <div className="flex items-center gap-3">
          <Image
            src={logo}
            alt="TourVista logo"
            width={50}
            height={50}
            className="w-[50px]"
            priority
          />
          <p className="font-bold text-xl text-gray-800">
            Tour<span className="text-orange-500">Vista</span>
          </p>
        </div>

        {/* Desktop Navigation Links */}
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

        {/* Desktop Auth Section */}
        <div className="hidden md:flex items-center gap-4">
          {!user ? (
            <>
              <button
                onClick={() => setShowLogin(true)}
                className="bg-orange-500 hover:bg-orange-600 px-6 py-2 rounded-lg text-white font-medium transition-all shadow-sm hover:shadow-md"
              >
                Login
              </button>
              <button
                onClick={() => setShowSignup(true)}
                className="bg-white border border-orange-400 hover:bg-orange-100 text-orange-500 px-6 py-2 rounded-lg font-medium transition-all"
              >
                Signup
              </button>
            </>
          ) : (
            <>
              <div className="bg-orange-500 text-white w-10 h-10 rounded-full flex items-center justify-center font-bold text-lg shadow">
                {user.name.charAt(0).toUpperCase()}
              </div>
              <button
                onClick={handleLogout}
                className="text-gray-700 hover:text-orange-500 text-sm font-medium"
              >
                Logout
              </button>
            </>
          )}
        </div>

        {/* Mobile Menu Button */}
        <div className="md:hidden flex items-center">
          <button
            onClick={() => setMobileMenuOpen(!isMobileMenuOpen)}
            className="p-2 text-gray-700 hover:text-orange-500"
          >
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              {isMobileMenuOpen ? (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M6 18L18 6M6 6l12 12"
                />
              ) : (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4 6h16M4 12h16M4 18h16"
                />
              )}
            </svg>
          </button>
        </div>
      </div>

      {/* Mobile Menu */}
      {isMobileMenuOpen && (
        <div className="absolute top-full left-0 right-0 bg-white p-5 shadow-lg md:hidden border-t border-gray-100 animate-fade-in">
          <ul className="flex flex-col gap-4">
            {NavbarLinks.map((navLink) => (
              <li key={navLink.id} className="border-b border-gray-100 pb-2">
                <Link
                  href={navLink.link}
                  onClick={() => setMobileMenuOpen(false)}
                  className={`block text-lg font-medium ${
                    isActiveLink(navLink.link)
                      ? "text-orange-500"
                      : "text-gray-800 hover:text-orange-500"
                  }`}
                >
                  {navLink.title}
                </Link>
              </li>
            ))}

            {!user ? (
              <li className="pt-2 space-y-2">
                <button
                  onClick={() => {
                    setShowLogin(true);
                    setMobileMenuOpen(false);
                  }}
                  className="w-full bg-orange-500 hover:bg-orange-600 px-6 py-3 rounded-lg text-white font-semibold text-center shadow-md hover:shadow-lg transition-all"
                >
                  Login
                </button>
                <button
                  onClick={() => {
                    setShowSignup(true);
                    setMobileMenuOpen(false);
                  }}
                  className="w-full bg-white hover:bg-orange-100 border border-orange-300 px-6 py-3 rounded-lg text-orange-500 font-semibold text-center shadow-sm hover:shadow-md transition-all"
                >
                  Signup
                </button>
              </li>
            ) : (
              <li className="flex items-center gap-4 pt-4">
                <div className="bg-orange-500 text-white w-10 h-10 rounded-full flex items-center justify-center font-bold text-lg shadow">
                  {user.name.charAt(0).toUpperCase()}
                </div>
                <button
                  onClick={() => {
                    handleLogout();
                    setMobileMenuOpen(false);
                  }}
                  className="text-gray-700 hover:text-orange-500 text-sm font-medium"
                >
                  Logout
                </button>
              </li>
            )}
          </ul>
        </div>
      )}

      {/* Modals */}
      <Modal
        title="Login to your account"
        isOpen={showLogin}
        onClose={() => setShowLogin(false)}
      >
        <LoginForm onSuccess={handleLoginSuccess} />
      </Modal>

      <Modal
        title="Create a new account"
        isOpen={showSignup}
        onClose={() => setShowSignup(false)}
      >
        <SignupForm onSuccess={handleSignupSuccess} />
      </Modal>
    </>
  );
};

export default Navbar;
