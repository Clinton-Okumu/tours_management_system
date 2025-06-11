const TourCard = ({ tour }) => {
  return (
    <div className="bg-base rounded-lg overflow-hidden hover:transform hover:scale-105 transition-all duration-300">
      <div
        className="h-48 bg-cover bg-center"
        style={{ backgroundImage: `url(${tour.image})` }}
      ></div>
      <div className="p-6">
        <h3 className="text-xl font-semibold mb-3">{tour.title}</h3>
        <p className="text-surface mb-4">{tour.description}</p>
        <button className="text-primary-light hover:text-primary font-semibold">
          Learn More →
        </button>
      </div>
    </div>
  );
};

export default TourCard;
