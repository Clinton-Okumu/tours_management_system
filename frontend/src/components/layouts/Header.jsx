import React from 'react';
import { MapPin } from 'lucide-react';
import { Link } from 'react-router-dom'; // Import Link

const Header = () => {
  return (
    <nav className="flex items-center justify-between px-6 py-4">
      <div className="flex items-center space-x-2">
        <MapPin className="w-6 h-6" />
        <Link to="/" className="text-xl font-semibold hover:text-primary-light transition-colors">
          Clint Tours
        </Link>
      </div>
      <div className="flex items-center space-x-8">
        <Link to="/tours" className="hover:text-primary-light transition-colors">Tours</Link>
        <Link to="/destinations" className="hover:text-primary-light transition-colors">Destinations</Link>
        <Link to="/about" className="hover:text-primary-light transition-colors">About Us</Link>
        <Link to="/contact" className="hover:text-primary-light transition-colors">Contact</Link>
        <Link to="/book" className="bg-primary hover:bg-primary-dark px-6 py-2 rounded-full transition-colors">
          Book Now
        </Link>
      </div>
    </nav>
  );
};

export default Header;
