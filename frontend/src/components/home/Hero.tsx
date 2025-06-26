import hero from "@/assets/hero.png";
import Image from "next/image";

const HeroImage = () => {
  return (
    <div className="relative w-full h-full overflow-hidden">
      <Image
        src={hero}
        alt="Beautiful travel destination"
        fill
        className="object-cover"
        priority
      />
      {/* Optional overlay for better text contrast */}
      <div className="absolute inset-0 bg-black/20"></div>
    </div>
  );
};

export default HeroImage;
