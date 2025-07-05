const ContactMap = () => {
  return (
    <div className="w-full max-w-4xl mx-auto mt-10">
      <iframe
        src="https://www.google.com/maps/embed?pb=..."
        className="w-full h-64 rounded-lg border-0"
        allowFullScreen
        loading="lazy"
        referrerPolicy="no-referrer-when-downgrade"
      />
    </div>
  );
};

export default ContactMap;
