import Image from "next/image";

import gallery01 from "@/assets/images/gallery-01.jpg";
import gallery02 from "@/assets/images/gallery-02.jpg";
import gallery03 from "@/assets/images/gallery-03.jpg";
import gallery04 from "@/assets/images/gallery-04.jpg";
import gallery05 from "@/assets/images/gallery-05.jpg";
import gallery06 from "@/assets/images/gallery-06.jpg";

const galleryImages = [
  { src: gallery01, alt: "Gallery 1" },
  { src: gallery02, alt: "Gallery 2" },
  { src: gallery03, alt: "Gallery 3" },
  { src: gallery04, alt: "Gallery 4" },
  { src: gallery05, alt: "Gallery 5" },
  { src: gallery06, alt: "Gallery 6" },
];

const GallerySection = () => {
  return (
    <section className="w-full py-20 px-4 md:px-12 lg:px-24 bg-gray-50">
      <h2 className="text-center text-2xl md:text-4xl font-bold mb-10 text-gray-800 tracking-wide">
        — OUR GALLERY —
      </h2>

      <div className="grid grid-cols-2 md:grid-cols-4 gap-4 auto-rows-[200px]">
        {galleryImages.map((img, index) => (
          <div
            key={index}
            className={`relative rounded-2xl overflow-hidden group shadow-md hover:shadow-xl transition-all duration-300 ${
              index % 3 === 0 ? "row-span-2" : "row-span-1"
            }`}
          >
            <Image
              src={img.src}
              alt={img.alt}
              className="object-cover w-full h-full group-hover:scale-105 transition-transform duration-300"
              placeholder="blur"
              fill
              sizes="(max-width: 768px) 100vw, 33vw"
            />
          </div>
        ))}
      </div>
    </section>
  );
};

export default GallerySection;
