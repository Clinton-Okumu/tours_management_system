import Image from "next/image";
import aboutimage from "@/assets/aboutimage1.jpg";

const AboutBody = () => {
  return (
    <section className="py-20 px-6 bg-white">
      <div className="max-w-6xl mx-auto flex flex-col md:flex-row items-center gap-12">
        {/* Image Left */}
        <div className="w-full md:w-1/2">
          <Image
            src={aboutimage}
            alt="About our tours"
            width={600}
            height={400}
            className="rounded-xl shadow-lg object-cover w-full h-auto"
            priority // Optional: improve LCP for above-the-fold content
          />
        </div>

        {/* Text Right */}
        <div className="w-full md:w-1/2 text-gray-700">
          <h2 className="text-4xl font-bold mb-6 text-gray-800">
            Discover the Heart of Kenya
          </h2>
          <p className="mb-4 leading-relaxed">
            At our core, we’re storytellers, adventurers, and conservationists
            dedicated to crafting journeys that connect travelers with the soul
            of Kenya.
          </p>
          <p className="mb-4 leading-relaxed">
            Whether it's tracking the Big Five across the savanna, enjoying
            cultural exchanges in Maasai villages, or exploring hidden gems off
            the beaten path, we ensure every moment is unforgettable.
          </p>
          <p className="leading-relaxed">
            With a passionate local team and a commitment to sustainability, our
            tours go beyond sightseeing — they immerse you in experiences that
            matter.
          </p>
        </div>
      </div>
    </section>
  );
};

export default AboutBody;
