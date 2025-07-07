import ContactDetails from "@/components/contact/contactdetails";
import ContactForm from "@/components/contact/contactform";
import ContactMap from "@/components/contact/contactmap";
import ContactHeader from "@/components/contact/header";

export default function Contact() {
  return (
    <main className="bg-white min-h-screen py-16 px-6 md:px-16 lg:px-32 space-y-24 shadow-xl rounded-xl">
      <ContactHeader />
      <ContactForm />
      <ContactDetails />
      <ContactMap />
    </main>
  );
}
