import Navbar from "../components/Navbar"
import AddPet from "../components/AddPet";
import PetSearchAndResults from "../components/searchPets";

function Home(){
    
    return (
        <>
        <Navbar/>
        <AddPet/>
        <PetSearchAndResults/>
        </>
    )
}

export default Home;