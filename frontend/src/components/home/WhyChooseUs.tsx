import { Shield, Star, MapPin, Clock, Users, Award } from "lucide-react";

const features = [
  {
    id: 1,
    icon: Shield,
    title: "Safe & Secure",
    description:
      "Your safety is our priority with certified guides and secure transportation.",
  },
  {
    id: 2,
    icon: Star,
    title: "5-Star Experience",
    description:
      "Premium quality tours with attention to every detail for unforgettable memories.",
  },
  {
    id: 3,
    icon: MapPin,
    title: "Local Expertise",
    description:
      "Deep local knowledge and exclusive access to hidden gems and authentic experiences.",
  },
  {
    id: 4,
    icon: Clock,
    title: "24/7 Support",
    description:
      "Round-the-clock customer support to assist you throughout your journey.",
  },
  {
    id: 5,
    icon: Users,
    title: "Small Groups",
    description:
      "Intimate group sizes for personalized attention and better experiences.",
  },
  {
    id: 6,
    icon: Award,
    title: "Award Winning",
    description:
      "Recognized excellence in tourism with multiple industry awards and certifications.",
  },
];

const WhyChooseUs = () => {
  return (
    <section className="w-full py-20 px-4 md:px-12 lg:px-24 bg-white">
      <div className="max-w-7xl mx-auto">
        <h2 className="text-center text-2xl md:text-4xl font-bold mb-4 text-gray-800 tracking-wide">
          — WHY CHOOSE US? —
        </h2>
        <p className="text-center text-gray-600 mb-16 max-w-2xl mx-auto">
          We're committed to providing exceptional travel experiences that
          exceed your expectations
        </p>

        <div className="grid gap-8 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
          {features.map((feature) => {
            const IconComponent = feature.icon;
            return (
              <div
                key={feature.id}
                className="text-center p-6 rounded-xl shadow-md hover:shadow-xl transition-shadow duration-300 group bg-white border border-gray-100"
              >
                <div className="inline-flex items-center justify-center w-16 h-16 bg-orange-100 rounded-full mb-6 group-hover:bg-orange-500 transition-colors duration-300">
                  <IconComponent className="w-8 h-8 text-orange-500 group-hover:text-white transition-colors duration-300" />
                </div>
                <h3 className="text-xl font-semibold text-gray-800 mb-3">
                  {feature.title}
                </h3>
                <p className="text-gray-600 leading-relaxed">
                  {feature.description}
                </p>
              </div>
            );
          })}
        </div>
      </div>
    </section>
  );
};

export default WhyChooseUs;
