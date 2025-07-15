"use client";

import { useState } from "react";
import { Star, MapPin, Calendar, Users, Heart } from "lucide-react";

interface TourCardProps {
  title: string;
  image: string;
  duration: number;
  difficulty: string;
  description: string;
}

export default function TourCard({
  title,
  image,
  duration,
  difficulty,
  description,
}: TourCardProps) {
  const [liked, setLiked] = useState(false);

  return (
    <div className="rounded-xl shadow-lg overflow-hidden bg-white relative hover:shadow-xl transition-shadow">
      <img src={image} alt={title} className="w-full h-48 object-cover" />
      <div className="p-4">
        <button
          onClick={() => setLiked(!liked)}
          className="absolute top-4 right-4 text-red-500 hover:scale-110 transition-transform"
          aria-label="Like tour"
        >
          <Heart fill={liked ? "#ef4444" : "none"} className="w-5 h-5" />
        </button>

        <h3 className="text-lg font-semibold text-gray-800 mb-1">{title}</h3>
        <p className="text-sm text-gray-600 line-clamp-3">{description}</p>

        <div className="mt-3 text-sm flex justify-between text-gray-700">
          <span className="flex items-center gap-1">
            <Calendar size={16} />
            {duration} days
          </span>
          <span className="flex items-center gap-1">
            <Users size={16} />
            {difficulty}
          </span>
        </div>
      </div>
    </div>
  );
}
