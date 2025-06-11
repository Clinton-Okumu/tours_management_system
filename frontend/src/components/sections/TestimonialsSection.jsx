import TestimonialCard from "../cards/TestimonialCard"; // Corrected path
import testimonials from "../../data/testimonials"; // Assuming you move data here

const TestimonialsSection = () => {
  // The testimonials data is now imported from src/data/testimonials.js
  return (
    <div className="px-6 py-16">
      <h2 className="text-3xl font-bold mb-12">Customer Testimonials</h2>
      <div className="space-y-8">
        {testimonials.map((testimonial) => (
          <TestimonialCard key={testimonial.id} testimonial={testimonial} />
        ))}
      </div>
    </div>
  );
};

export default TestimonialsSection;
