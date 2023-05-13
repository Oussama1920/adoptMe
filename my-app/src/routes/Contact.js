import Navbar from "../components/Navbar"
import Hero from "../components/Hero";
import AboutImg from "../assets/dogContact.jpg"
import ContactForm from "../components/ContactForm";
function Contact(){
    return (
        <>
        <Navbar/>
        <Hero 
         cName="hero"
         heroImg={AboutImg}
         title="Contact"
         btnClass="hide"
        />
        <ContactForm/>
        </>
    )
}

export default Contact;