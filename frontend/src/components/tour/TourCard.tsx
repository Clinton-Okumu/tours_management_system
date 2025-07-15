import Image from "next/image";
import Link from "next/link";

type TourCardProps = {
  id: number;
  title: string;
  price: number;
  duration: string;
  image: string;
};

export default function TourCard({
  id,
  title,
  price,
  duration,
  image,
}: TourCardProps) {
  return (
    <div className="rounded-xl overflow-hidden shadow-md bg-white">
      <Image
        src={image}
        alt={title}
        width={400}
        height={250}
        className="w-full h-60 object-cover"
      />
      <div className="p-4">
        <h3 className="text-xl font-semibold">{title}</h3>
        <p className="text-sm text-gray-600">{duration}</p>
        <p className="text-md font-bold text-green-600">${price}</p>
        <Link href={`/tours/${id}`}>
          <button className="mt-3 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
            View Details
          </button>
        </Link>
      </div>
    </div>
  );
}
