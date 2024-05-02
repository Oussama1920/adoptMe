import Navbar from "../components/Navbar"
import AddPet from "../components/AddPet";
import PetSearchAndResults from "../components/searchPets";

function Home(){
    
    return (
        <>
        <Navbar/>
        <PetSearchAndResults/>
        </>
    )
}

export default Home;