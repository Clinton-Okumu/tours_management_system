import BookingForm from "@/components/home/BookingForm";
import OurDestinations from "@/components/home/FeaturedTours";
import GallerySection from "@/components/home/Gallery";
import HeroImage from "@/components/home/Hero";
import WhyChooseUs from "@/components/home/WhyChooseUs";
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

      {/* Gallery section */}
      <GallerySection />
    </>
  );
}
