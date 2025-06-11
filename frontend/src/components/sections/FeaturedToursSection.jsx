import React from "react";
import TourCard from "../cards/TourCard";
import featuredTours from "../../data/featuredTours";

const FeaturedToursSection = () => {
  return (
    <div className="px-6 py-16">
      <h2 className="text-3xl font-bold mb-12">Featured Tours</h2>
      <div className="grid md:grid-cols-3 gap-8">
        {featuredTours.map((tour) => (
          <TourCard key={tour.id} tour={tour} />
        ))}
      </div>
    </div>
  );
};

export default FeaturedToursSection;
