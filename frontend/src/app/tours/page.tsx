import TourCard from "@/components/tour/TourCard";
import TourHeader from "@/components/tour/TourHeader";
import { tours } from "@/data/tours";

export default function ToursPage() {
  return (
    <main className="min-h-screen bg-white px-4 py-8">
      <TourHeader />
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 mt-6">
        {tours.map((tour) => (
          <TourCard key={tour.id} {...tour} />
        ))}
      </div>
    </main>
  );
}
