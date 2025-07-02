import image from "@/assets/image1.jpg";
import image1 from "@/assets/image2.jpg";
import image2 from "@/assets/image3.jpg";
import image3 from "@/assets/image4.jpg";
import Image from "next/image";

const team = [
  {
    name: "Clinton Okumu",
    role: "Founder & CEO",
    bio: "Visionary leader with 15+ years of tourism experience, passionate about connecting the world to Kenya.",
    image: image,
  },
  {
    name: "James Mwangi",
    role: "Head of Operations",
    bio: "Former KWS ranger turned logistics expert, ensuring every safari is seamless and safe.",
    image: image1,
  },
  {
    name: "Amina Hassan",
    role: "Guest Experience Manager",
    bio: "Fluent in 5 languages, Amina ensures every traveler feels welcome from arrival to goodbye.",
    image: image2,
  },
  {
    name: "David Kipchoge",
    role: "Senior Safari Guide",
    bio: "Third-generation Maasai guide with deep knowledge of Kenya’s wildlife and storytelling.",
    image: image3,
  },
];

const TeamSection = () => {
  return (
    <section className="py-20 px-6 bg-gray-50">
      <div className="max-w-6xl mx-auto text-center">
        <h2 className="text-4xl md:text-5xl font-bold text-gray-800 mb-6">
          - Meet Our Team -
        </h2>
        <p className="text-xl text-gray-600 mb-12">
          The passionate people behind your unforgettable experience
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {team.map((member, index) => (
            <div
              key={index}
              className="bg-white rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105"
            >
              <div className="relative w-24 h-24 mx-auto mb-6 rounded-full overflow-hidden shadow-md">
                <Image
                  src={member.image}
                  alt={member.name}
                  fill
                  className="object-cover"
                />
              </div>
              <h3 className="text-xl font-bold text-gray-800 mb-1">
                {member.name}
              </h3>
              <p className="text-orange-600 font-medium mb-3">{member.role}</p>
              <p className="text-sm text-gray-600">{member.bio}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default TeamSection;
