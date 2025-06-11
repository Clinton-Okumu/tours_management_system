import React from "react";
import { Star } from "lucide-react";

const StarRating = ({ rating }) => {
  return (
    <div className="flex items-center space-x-1">
      {Array.from({ length: 5 }, (_, index) => (
        <Star
          key={index}
          className={`w-4 h-4 ${index < rating ? "fill-primary-light text-primary-light" : "text-muted"}`}
        />
      ))}
    </div>
  );
};

export default StarRating;
