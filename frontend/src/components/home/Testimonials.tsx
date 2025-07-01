import { Quote, Star } from "lucide-react";

const testimonials = [
  {
    id: 1,
    name: "Sarah Johnson",
    location: "London, UK",
    rating: 5,
    text: "An absolutely incredible experience! The Maasai Mara safari exceeded all our expectations. Our guide was knowledgeable and the accommodations were perfect.",
    date: "March 2024",
  },
  {
    id: 2,
    name: "Michael Chen",
    location: "Toronto, Canada",
    rating: 5,
    text: "The Mount Kenya trek was challenging but so rewarding. The team took excellent care of us and the views were breathtaking. Highly recommended!",
    date: "January 2024",
  },
  {
    id: 3,
    name: "Emma Rodriguez",
    location: "Madrid, Spain",
    rating: 5,
    text: "Diani Beach was paradise! Everything was organized perfectly from airport pickup to the beautiful beachfront resort. Can't wait to return!",
    date: "February 2024",
  },
  {
    id: 4,
    name: "David Thompson",
    location: "Sydney, Australia",
    rating: 5,
    text: "Professional service from start to finish. The cultural experiences were authentic and the wildlife encounters were unforgettable. Thank you for an amazing trip!",
    date: "December 2023",
  },
  {
    id: 5,
    name: "Lisa Park",
    location: "Seoul, South Korea",
    rating: 5,
    text: "The attention to detail was impressive. Every aspect of our Kenya tour was well-planned and executed flawlessly. The memories will last a lifetime!",
    date: "November 2023",
  },
  {
    id: 6,
    name: "James Wilson",
    location: "New York, USA",
    rating: 5,
    text: "Outstanding customer service and incredible destinations. The small group size made it feel very personal and exclusive. Worth every penny!",
    date: "October 2023",
  },
];

const Testimonials = () => {
  return (
    <section className="w-full py-20 px-4 md:px-12 lg:px-24 bg-gray-50">
      <div className="max-w-7xl mx-auto">
        <h2 className="text-center text-2xl md:text-4xl font-bold mb-4 text-gray-800 tracking-wide">
          — TESTIMONIALS —
        </h2>
        <p className="text-center text-gray-600 mb-16 max-w-2xl mx-auto">
          Don't just take our word for it - hear what our travelers have to say
          about their experiences
        </p>

        <div className="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
          {testimonials.map((testimonial) => (
            <div
              key={testimonial.id}
              className="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300 relative"
            >
              <Quote className="absolute top-4 right-4 w-8 h-8 text-orange-200" />

              <div className="flex items-center mb-4">
                {[...Array(testimonial.rating)].map((_, i) => (
                  <Star
                    key={i}
                    className="w-5 h-5 text-yellow-400 fill-current"
                  />
                ))}
              </div>

              <p className="text-gray-700 mb-6 leading-relaxed italic">
                "{testimonial.text}"
              </p>

              <div className="border-t pt-4">
                <h4 className="font-semibold text-gray-800">
                  {testimonial.name}
                </h4>
                <p className="text-sm text-gray-500">{testimonial.location}</p>
                <p className="text-xs text-gray-400 mt-1">{testimonial.date}</p>
              </div>
            </div>
          ))}
        </div>

        <div className="text-center mt-12">
          <div className="inline-flex items-center space-x-2 bg-orange-100 px-6 py-3 rounded-full mb-6">
            <div className="flex">
              {[...Array(5)].map((_, i) => (
                <Star
                  key={i}
                  className="w-5 h-5 text-yellow-400 fill-current"
                />
              ))}
            </div>
            <span className="text-gray-700 font-semibold">4.9/5</span>
            <span className="text-gray-600">from 500+ reviews</span>
          </div>

          <button className="bg-orange-500 hover:bg-orange-600 text-white font-semibold px-8 py-3 rounded-lg transition-colors duration-200 shadow-md hover:shadow-lg">
            Write a Review
          </button>
        </div>
      </div>
    </section>
  );
};

export default Testimonials;
