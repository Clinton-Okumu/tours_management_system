"use client";
import React from "react";
import Image from "next/image";
import { Facebook, Twitter, Instagram, Mail } from "lucide-react";
import logo from "@/assets/logo.png";

const Footer = () => {
  return (
    <footer className="bg-gray-900 text-gray-300 pt-12 pb-6">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Top Section */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 border-b border-gray-700 pb-10">
          {/* Logo */}
          <div>
            <div className="flex items-center gap-3">
              <Image src={logo} alt="TourVista logo" width={40} height={40} />
              <span className="text-xl font-bold text-white">
                Tour<span className="text-orange-500">Vista</span>
              </span>
            </div>
            <p className="mt-4 text-sm text-gray-400">
              Explore the world with confidence. Book your next adventure with
              TourVista.
            </p>
          </div>

          {/* Links */}
          <div>
            <h4 className="text-white font-semibold mb-4">Company</h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a href="#" className="hover:text-orange-500">
                  About Us
                </a>
              </li>
              <li>
                <a href="#" className="hover:text-orange-500">
                  Careers
                </a>
              </li>
              <li>
                <a href="#" className="hover:text-orange-500">
                  Blog
                </a>
              </li>
            </ul>
          </div>

          <div>
            <h4 className="text-white font-semibold mb-4">Support</h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a href="#" className="hover:text-orange-500">
                  Help Center
                </a>
              </li>
              <li>
                <a href="#" className="hover:text-orange-500">
                  Contact
                </a>
              </li>
              <li>
                <a href="#" className="hover:text-orange-500">
                  Terms & Privacy
                </a>
              </li>
            </ul>
          </div>

          {/* Newsletter */}
          <div>
            <h4 className="text-white font-semibold mb-4">Subscribe</h4>
            <p className="text-sm text-gray-400 mb-4">
              Get updates on new tours and offers.
            </p>
            <form className="flex flex-col sm:flex-row gap-2">
              <input
                type="email"
                placeholder="Enter your email"
                className="px-4 py-2 rounded-md bg-gray-800 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-orange-500"
              />
              <button
                type="submit"
                className="px-4 py-2 bg-orange-500 hover:bg-orange-600 text-white rounded-md transition"
              >
                Subscribe
              </button>
            </form>
          </div>
        </div>

        {/* Bottom Section */}
        <div className="flex flex-col sm:flex-row justify-between items-center mt-8 text-sm text-gray-500">
          <p>
            &copy; {new Date().getFullYear()} TourVista. All rights reserved.
          </p>
          <div className="flex gap-4 mt-4 sm:mt-0">
            <a href="#" className="hover:text-orange-500">
              <Facebook size={20} />
            </a>
            <a href="#" className="hover:text-orange-500">
              <Twitter size={20} />
            </a>
            <a href="#" className="hover:text-orange-500">
              <Instagram size={20} />
            </a>
            <a
              href="mailto:hello@tourvista.com"
              className="hover:text-orange-500"
            >
              <Mail size={20} />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
