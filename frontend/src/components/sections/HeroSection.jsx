const HeroSection = () => {
  return (
    <div className="relative px-6 py-20">
      <div
        className="relative h-96 rounded-2xl overflow-hidden bg-cover bg-center"
        style={{
          backgroundImage:
            "url('https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=1200&h=600&fit=crop')",
        }}
      >
        <div className="absolute inset-0 bg-black bg-opacity-30"></div>
        <div className="relative z-10 flex flex-col items-center justify-center h-full text-center px-4">
          <h1 className="text-5xl font-bold mb-4">
            Explore the World with Clint Tours
          </h1>
          <p className="text-lg mb-8 max-w-2xl">
            Discover unforgettable journeys with our expertly curated tours.
            From breathtaking landscapes to vibrant cityscapes, we have the
            perfect adventure for you.
          </p>
          <button className="bg-primary hover:bg-primary-dark px-8 py-3 rounded-full text-lg font-semibold transition-colors">
            View Tours
          </button>
        </div>
      </div>
    </div>
  );
};

export default HeroSection;
