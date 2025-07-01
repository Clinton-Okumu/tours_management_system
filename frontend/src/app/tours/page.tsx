import { 
  Heart, 
  Users, 
  Globe, 
  Award, 
  Camera, 
  MapPin, 
  Star, 
  Calendar,
  CheckCircle,
  ArrowRight,
  Mail,
  Phone,
  Instagram,
  Facebook,
  Twitter
} from "lucide-react";
import { useState } from "react";

const AboutPage = () => {
  const [activeYear, setActiveYear] = useState(2020);

  const timeline = [
    {
      year: 2018,
      title: "The Dream Begins",
      description: "Founded by passionate travelers who fell in love with Kenya's incredible wildlife and warm hospitality. Started with a single safari vehicle and a big dream.",
      milestone: "First Safari"
    },
    {
      year: 2019,
      title: "Growing Community",
      description: "Expanded our team with local Maasai guides and conservation experts. Launched our first community-based tourism initiatives.",
      milestone: "100+ Travelers"
    },
    {
      year: 2020,
      title: "Adapting & Innovating",
      description: "Pivoted during global challenges by creating virtual safari experiences and supporting local communities through difficult times.",
      milestone: "Community Support"
    },
    {
      year: 2021,
      title: "Award Recognition",
      description: "Received 'Best Sustainable Tourism Operator' award. Launched our wildlife conservation partnership program.",
      milestone: "1000+ Travelers"
    },
    {
      year: 2022,
      title: "Digital Excellence",
      description: "Launched our mobile app and introduced AI-powered trip planning. Expanded to serve travelers from 50+ countries.",
      milestone: "50+ Countries"
    },
    {
      year: 2023,
      title: "Conservation Impact",
      description: "Established the Kenya Wildlife Foundation. Our travelers contributed to protecting over 10,000 acres of wildlife habitat.",
      milestone: "3000+ Travelers"
    },
    {
      year: 2024,
      title: "Global Recognition",
      description: "Featured in National Geographic. Reached 5000+ travelers served. Launched luxury eco-lodge partnerships.",
      milestone: "5000+ Travelers"
    }
  ];

  const team = [
    {
      name: "Sarah Kimani",
      role: "Founder & CEO",
      bio: "Born in Nairobi, Sarah combines her business degree with 15+ years of tourism experience. Her vision: making Kenya accessible to the world while preserving its natural beauty.",
      image: "SK",
      specialties: ["Strategic Planning", "Sustainable Tourism", "Community Relations"]
    },
    {
      name: "James Mwangi",
      role: "Head of Operations",
      bio: "Former Kenya Wildlife Service ranger turned operations expert. James ensures every safari runs smoothly while maintaining the highest safety standards.",
      image: "JM",
      specialties: ["Wildlife Expertise", "Safety Management", "Logistics"]
    },
    {
      name: "Amina Hassan",
      role: "Guest Experience Manager",
      bio: "Fluent in 5 languages, Amina's hospitality background ensures every traveler feels welcomed from arrival to departure. She's your go-to for special requests.",
      image: "AH",
      specialties: ["Customer Service", "Cultural Exchange", "Languages"]
    },
    {
      name: "David Kipchoge",
      role: "Senior Safari Guide",
      bio: "Third-generation Maasai guide with unmatched knowledge of Kenya's ecosystems. David's storytelling brings the savanna to life for every guest.",
      image: "DK",
      specialties: ["Wildlife Tracking", "Cultural Heritage", "Photography"]
    }
  ];

  const stats = [
    { number: "5,000+", label: "Happy Travelers", icon: Users },
    { number: "4.9/5", label: "Average Rating", icon: Star },
    { number: "50+", label: "Countries Served", icon: Globe },
    { number: "98%", label: "Return Rate", icon: Heart },
    { number: "15+", label: "Conservation Projects", icon: Award },
    { number: "24/7", label: "Support Available", icon: Phone }
  ];

  const values = [
    {
      title: "Authentic Experiences",
      description: "We believe in genuine cultural exchange and real connections with local communities.",
      icon: Heart
    },
    {
      title: "Conservation First",
      description: "Every tour contributes to wildlife protection and habitat preservation.",
      icon: Globe
    },
    {
      title: "Safety Excellence",
      description: "Your security and well-being are our top priorities on every adventure.",
      icon: CheckCircle
    },
    {
      title: "Local Expertise",
      description: "Our guides are local experts who share their deep knowledge and passion.",
      icon: Users
    }
  ];

  return (
    <div className="min-h-screen bg-white">
      {/* Hero Section */}
      <section className="relative min-h-screen flex items-center justify-center overflow-hidden">
        {/* Background with gradient overlay */}
        <div className="absolute inset-0 bg-gradient-to-br from-orange-900 via-orange-800 to-yellow-700"></div>
        <div className="absolute inset-0 bg-black opacity-40"></div>
        
        {/* Animated background elements */}
        <div className="absolute inset-0 overflow-hidden">
          <div className="absolute top-20 left-20 w-64 h-64 bg-orange-400 rounded-full opacity-20 blur-3xl animate-pulse"></div>
          <div className="absolute bottom-20 right-20 w-80 h-80 bg-yellow-400 rounded-full opacity-20 blur-3xl animate-pulse delay-1000"></div>
          <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-orange-300 rounded-full opacity-10 blur-3xl animate-pulse delay-2000"></div>
        </div>

        <div className="relative z-10 text-center text-white max-w-4xl mx-auto px-6">
          <h1 className="text-6xl md:text-8xl font-bold mb-6 leading-tight">
            About
            <span className="block bg-gradient-to-r from-orange-400 to-yellow-400 bg-clip-text text-transparent">
              Our Story
            </span>
          </h1>
          <p className="text-xl md:text-2xl mb-8 opacity-90 leading-relaxed">
            Our journey to make your journeys unforgettable
          </p>
          <div className="w-24 h-1 bg-gradient-to-r from-orange-400 to-yellow-400 mx-auto mb-8"></div>
          <p className="text-lg opacity-80 max-w-2xl mx-auto leading-relaxed">
            Born from a love for Kenya's wild beauty and rich culture, we've spent years crafting 
            authentic adventures that connect travelers with the heart of East Africa.
          </p>
        </div>

        {/* Scroll indicator */}
        <div className="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce">
          <div className="w-6 h-10 border-2 border-white rounded-full flex justify-center">
            <div className="w-1 h-3 bg-white rounded-full mt-2 animate-pulse"></div>
          </div>
        </div>
      </section>

      {/* Company Story Timeline */}
      <section className="py-20 px-6 bg-gray-50">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
              Our Journey
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              From humble beginnings to becoming Kenya's most trusted travel partner
            </p>
          </div>

          <div className="relative">
            {/* Timeline line */}
            <div className="absolute left-1/2 transform -translate-x-1/2 w-1 h-full bg-gradient-to-b from-orange-500 to-yellow-500 rounded-full"></div>

            {timeline.map((item, index) => (
              <div
                key={item.year}
                className={relative flex items-center mb-12 ${
                  index % 2 === 0 ? 'flex-row' : 'flex-row-reverse'
                }}
                onMouseEnter={() => setActiveYear(item.year)}
              >
                <div className={w-1/2 ${index % 2 === 0 ? 'pr-8 text-right' : 'pl-8 text-left'}}>
                  <div className={bg-white p-6 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105 ${
                    activeYear === item.year ? 'ring-2 ring-orange-500' : ''
                  }}>
                    <div className="flex items-center justify-between mb-3">
                      <span className="text-2xl font-bold text-orange-600">{item.year}</span>
                      <span className="bg-orange-100 text-orange-700 px-3 py-1 rounded-full text-sm font-medium">
                        {item.milestone}
                      </span>
                    </div>
                    <h3 className="text-xl font-bold text-gray-800 mb-2">{item.title}</h3>
                    <p className="text-gray-600 leading-relaxed">{item.description}</p>
                  </div>
                </div>

                {/* Timeline dot */}
                <div className="absolute left-1/2 transform -translate-x-1/2 w-6 h-6 bg-gradient-to-r from-orange-500 to-yellow-500 rounded-full border-4 border-white shadow-lg z-10"></div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Values Section */}
      <section className="py-20 px-6 bg-white">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
              Our Values
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              The principles that guide every adventure we create
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {values.map((value, index) => (
              <div
                key={index}
                className="text-center group hover:scale-105 transition-all duration-300"
              >
                <div className="w-20 h-20 bg-gradient-to-r from-orange-500 to-yellow-500 rounded-full flex items-center justify-center mx-auto mb-6 group-hover:shadow-xl transition-shadow duration-300">
                  <value.icon className="w-10 h-10 text-white" />
                </div>
                <h3 className="text-xl font-bold text-gray-800 mb-3">{value.title}</h3>
                <p className="text-gray-600 leading-relaxed">{value.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Meet the Team */}
      <section className="py-20 px-6 bg-gray-50">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
              Meet Our Team
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              The passionate people behind your unforgettable experiences
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {team.map((member, index) => (
              <div
                key={index}
                className="bg-white rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105 group"
              >
                <div className="w-24 h-24 bg-gradient-to-r from-orange-500 to-yellow-500 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-6 group-hover:scale-110 transition-transform duration-300">
                  {member.image}
                </div>
                <h3 className="text-xl font-bold text-gray-800 text-center mb-2">
                  {member.name}
                </h3>
                <p className="text-orange-600 text-center font-medium mb-4">
                  {member.role}
                </p>
                <p className="text-gray-600 text-sm leading-relaxed mb-4">
                  {member.bio}
                </p>
                <div className="flex flex-wrap gap-2 justify-center">
                  {member.specialties.map((specialty, idx) => (
                    <span
                      key={idx}
                      className="bg-orange-100 text-orange-700 px-2 py-1 rounded-full text-xs font-medium"
                    >
                      {specialty}
                    </span>
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-20 px-6 bg-gradient-to-r from-orange-600 to-yellow-600 text-white">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold mb-4">
              Our Impact
            </h2>
            <p className="text-xl opacity-90 max-w-3xl mx-auto">
              Numbers that tell the story of our commitment to excellence
            </p>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-8">
            {stats.map((stat, index) => (
              <div
                key={index}
                className="text-center group hover:scale-110 transition-transform duration-300"
              >
                <div className="w-16 h-16 bg-white bg-opacity-20 rounded-full flex items-center justify-center mx-auto mb-4 group-hover:bg-opacity-30 transition-all duration-300">
                  <stat.icon className="w-8 h-8" />
                </div>
                <div className="text-3xl md:text-4xl font-bold mb-2">
                  {stat.number}
                </div>
                <div className="text-sm opacity-90 font-medium">
                  {stat.label}
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Gallery/Culture Section */}
      <section className="py-20 px-6 bg-white">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
              Behind the Scenes
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              A glimpse into our culture, values, and the authentic Kenya we love to share
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[
              { title: "Team Planning Session", desc: "Crafting your perfect adventure" },
              { title: "Community Outreach", desc: "Supporting local Maasai villages" },
              { title: "Wildlife Conservation", desc: "Protecting Kenya's precious wildlife" },
              { title: "Guide Training", desc: "Continuous learning and improvement" },
              { title: "Cultural Exchange", desc: "Authentic interactions with locals" },
              { title: "Sustainable Practices", desc: "Eco-friendly tourism initiatives" }
            ].map((item, index) => (
              <div
                key={index}
                className="relative group overflow-hidden rounded-2xl bg-gradient-to-br from-orange-100 to-yellow-100 h-64 hover:shadow-xl transition-all duration-300"
              >
                <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent opacity-60"></div>
                <div className="absolute bottom-6 left-6 text-white">
                  <h3 className="text-xl font-bold mb-2">{item.title}</h3>
                  <p className="text-sm opacity-90">{item.desc}</p>
                </div>
                <div className="absolute top-6 right-6">
                  <Camera className="w-8 h-8 text-white opacity-70" />
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 px-6 bg-gray-900 text-white">
        <div className="max-w-4xl mx-auto text-center">
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Ready to Start Your
            <span className="block bg-gradient-to-r from-orange-400 to-yellow-400 bg-clip-text text-transparent">
              Kenya Adventure?
            </span>
          </h2>
          <p className="text-xl opacity-90 mb-8 max-w-2xl mx-auto leading-relaxed">
            Join thousands of travelers who've discovered the magic of Kenya with us. 
            Your unforgettable journey awaits.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
            <button className="bg-gradient-to-r from-orange-500 to-yellow-500 hover:from-orange-600 hover:to-yellow-600 text-white font-bold px-8 py-4 rounded-xl transition-all duration-300 shadow-lg hover:shadow-xl hover:scale-105 transform flex items-center justify-center">
              Explore Our Tours
              <ArrowRight className="ml-2 w-5 h-5" />
            </button>
            <button className="bg-transparent border-2 border-white hover:bg-white hover:text-gray-900 text-white font-semibold px-8 py-4 rounded-xl transition-all duration-300 flex items-center justify-center">
              Contact Us
              <Mail className="ml-2 w-5 h-5" />
            </button>
          </div>

          <div className="flex justify-center space-x-6">
            <a href="#" className="text-gray-400 hover:text-orange-400 transition-colors duration-300">
              <Instagram className="w-6 h-6" />
            </a>
            <a href="#" className="text-gray-400 hover:text-orange-400 transition-colors duration-300">
              <Facebook className="w-6 h-6" />
            </a>
            <a href="#" className="text-gray-400 hover:text-orange-400 transition-colors duration-300">
              <Twitter className="w-6 h-6" />
            </a>
          </div>
        </div>
      </section>
    </div>
  );
};

export default AboutPage;
