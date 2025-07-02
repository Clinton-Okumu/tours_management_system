import BookingForm from "@/components/home/BookingForm";
import HeroImage from "@/components/home/Hero";
import Testimonials from "@/components/home/Testimonials";
import WhyChooseUs from "@/components/home/WhyChooseUs";
import OurDestinations from "@/components/home/FeaturedTours";
export default function HomePage() {
  return (
    <>
      {/* Hero section with booking form overlay */}
      <section className="relative w-full h-[70vh] mb-16">
        <HeroImage />
        <BookingForm />
      </section>

      {/* Destination cards */}
      <OurDestinations />

      {/* why choose us section */}
      <WhyChooseUs />

      {/* testimonials section */}
      <Testimonials />
    </>
  );
}
