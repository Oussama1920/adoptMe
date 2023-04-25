import Navbar from "../components/Navbar"
import Hero from "../components/Hero";
import AboutImg from "../assets/pet1.jpg"
function About(){
    return (
        <>
        <Navbar/>
        <Hero 
         cName="hero-mid"
         heroImg={AboutImg}
         title="About"
         btnClass="hide"
        />        
        </>
    )
}

export default About;