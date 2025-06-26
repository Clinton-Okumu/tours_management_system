import pic1 from "@/assets/pic1.png";
import pic2 from "@/assets/pic2.png";
import pic3 from "@/assets/pic3.png";
import Image from "next/image";

const destinations = [
  {
    id: 1,
    name: "Diani Beach",
    description: "Pristine white sand and turquoise waters.",
    image: pic1,
  },
  {
    id: 2,
    name: "Maasai Mara",
    description: "Experience the great migration and wildlife.",
    image: pic2,
  },
  {
    id: 3,
    name: "Mount Kenya",
    description: "Adventure and stunning mountain views.",
    image: pic3,
  },
];

const OurDestinations = () => {
  return (
    <>
      <section className="w-full py-20 px-4 md:px-12 lg:px-24 bg-gray-50 ">
        <h2 className="text-center text-2xl md:text-4xl font-bold mb-10 text-gray-800 tracking-wide">
          — OUR DESTINATIONS —
        </h2>

        <div className="grid gap-8 grid-cols-1 md:grid-cols-3">
          {destinations.map((dest) => (
            <div
              key={dest.id}
              className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
            >
              <div className="h-56 relative">
                <Image
                  src={dest.image}
                  alt={dest.name}
                  fill
                  className="object-cover"
                />
              </div>
              <div className="p-5">
                <h3 className="text-xl font-semibold text-orange-500">
                  {dest.name}
                </h3>
                <p className="text-gray-600 mt-2 text-sm">{dest.description}</p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* More Destinations CTA */}
      <section className="text-center py-10 bg-gray-50">
        <button className="bg-orange-500 hover:bg-orange-600 text-white font-semibold px-6 py-3 rounded-lg transition">
          More Destinations &raquo;
        </button>
      </section>
    </>
  );
};
export default OurDestinations;
