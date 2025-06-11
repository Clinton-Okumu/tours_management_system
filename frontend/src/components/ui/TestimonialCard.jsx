import React from "react";
import { ThumbsUp, ThumbsDown } from "lucide-react";
import StarRating from "../common/StarRating"; // Corrected path

const TestimonialCard = ({ testimonial }) => {
  return (
    <div className="bg-base rounded-lg p-6">
      <div className="flex items-start space-x-4">
        <img
          src={testimonial.avatar}
          alt={testimonial.name}
          className="w-12 h-12 rounded-full object-cover"
        />
        <div className="flex-1">
          <div className="flex items-center justify-between mb-2">
            <div>
              <h4 className="font-semibold">{testimonial.name}</h4>
              <p className="text-muted text-sm">{testimonial.date}</p>
            </div>
          </div>
          <div className="mb-3">
            <StarRating rating={testimonial.rating} />
          </div>
          <p className="text-surface mb-4">{testimonial.review}</p>
          <div className="flex items-center space-x-4 text-sm text-muted">
            <div className="flex items-center space-x-1">
              <ThumbsUp className="w-4 h-4" />
              <span>{testimonial.likes}</span>
            </div>
            <div className="flex items-center space-x-1">
              <ThumbsDown className="w-4 h-4" />
              <span>{testimonial.dislikes}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TestimonialCard;
