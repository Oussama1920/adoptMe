import Hero from "../components/Hero";
import Navbar from "../components/Navbar"
import Results from "../components/Results";

function Home(){
    return (
        <>
        <Navbar/>
        <Hero 
         cName="hero"
         heroImg="https://wallpaperaccess.com/full/497375.jpg"
         title="Your Journey Your Story"
         text="Choose Your Favourite Destination."
         buttonText="Travel Plan"
         url ="/"
         btnClass="show"
        />
        <Results/>
        </>
    )
}

export default Home;