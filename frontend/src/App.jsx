import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

// Import Layout Component
import MainLayout from "./components/layout/MainLayout";

// Import Pages
import HomePage from "./pages/HomePage";
import ToursPage from "./pages/ToursPage";
import DestinationsPage from "./pages/DestinationsPage";
import AboutUsPage from "./pages/AboutUsPage";
import ContactPage from "./pages/ContactPage";
import NotFoundPage from "./pages/NotFoundPage"; // You'll create this

const App = () => {
  return (
    <Router>
      <MainLayout>
        {" "}
        {/* MainLayout wraps all routes */}
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/tours" element={<ToursPage />} />
          <Route path="/destinations" element={<DestinationsPage />} />
          <Route path="/about" element={<AboutUsPage />} />
          <Route path="/contact" element={<ContactPage />} />
          {/* Add more routes for specific tour details, booking page, privacy, terms etc. */}
          {/* Example for a dynamic route for a single tour */}
          <Route
            path="/tours/:id"
            element={<div>Tour Detail Page for ID</div>}
          />

          {/* Catch-all route for 404 Not Found */}
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </MainLayout>
    </Router>
  );
};

export default App;
