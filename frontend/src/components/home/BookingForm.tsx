const BookingForm = () => {
  return (
    <div className="absolute bottom-0 left-1/2 transform -translate-x-1/2 translate-y-1/2 w-full max-w-5xl px-4 z-10">
      <form className="bg-white p-6 md:p-8 rounded-xl shadow-2xl grid grid-cols-1 md:grid-cols-4 gap-4 border border-gray-100">
        <div>
          <label
            htmlFor="destination"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Destination
          </label>
          <input
            type="text"
            id="destination"
            placeholder="e.g. Nairobi"
            className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors"
          />
        </div>
        <div>
          <label
            htmlFor="date"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Date
          </label>
          <input
            type="date"
            id="date"
            className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors"
          />
        </div>
        <div>
          <label
            htmlFor="guests"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Guests
          </label>
          <input
            type="number"
            id="guests"
            min={1}
            placeholder="1"
            className="w-full border border-gray-300 rounded-lg px-4 py-3 focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-colors"
          />
        </div>
        <div className="flex items-end">
          <button
            type="submit"
            className="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 rounded-lg transition-colors duration-200 shadow-md hover:shadow-lg"
          >
            Check Availability
          </button>
        </div>
      </form>
    </div>
  );
};

export default BookingForm;
