import React from 'react';
import { Link } from 'react-router-dom';

const Footer = () => {
  return (
    <footer className="px-6 py-16 border-t border-surface-dark">
      <div className="flex flex-wrap justify-center space-x-8 mb-8 text-muted">
        <Link to="/about" className="hover:text-surface-light transition-colors">About Us</Link>
        <Link to="/contact" className="hover:text-surface-light transition-colors">Contact</Link>
        <Link to="/privacy" className="hover:text-surface-light transition-colors">Privacy Policy</Link>
        <Link to="/terms" className="hover:text-surface-light transition-colors">Terms of Service</Link>
      </div>
      <div className="text-center text-muted">
        <p>© 2023 Clint Tours. All rights reserved.</p>
      </div>
    </footer>
  );
};

export default Footer;
