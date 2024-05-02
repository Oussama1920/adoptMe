import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Navbar from "../components/Navbar"
import "./PetPageStyle.css";
import Pet from "./PetAnnonce";

const PetPage = () => {
  const { id } = useParams();
  const [pet, setPet] = useState(null);

  useEffect(() => {
    // Fetch pet details using the ID from the URL parameter
    fetch(`http://localhost:8080/v1/pets/pet/${id}`)
      .then(response => response.json())
      .then(data => setPet(data.pet))
      .catch(error => console.error('Error fetching pet details:', error));
  }, [id]);
  
  if (!pet) {
    return <div>Loading...</div>;
  }

  return (
    <div className="pet-page">
      <Navbar /> 
      <div className="pet-ad">
        <h1>{pet.name}</h1>
        <p>Type: {pet.type}</p>
        <p>Age: {pet.age}</p>
        <p>createdAt: {pet.created_at}</p>
        <div className="pet-images-grid">

        <Pet
                    type={pet.type}
                    name={pet.name}
                    age={pet.age}
                    available={pet.available}
                    images={pet.images}
                    id={pet.id}
        /> 
        </div>   
        <div className="pet-images-grid">
          {pet.images && pet.images.length > 0 ? (
            pet.images.map((image, index) => (
              <img key={index} src={image.data_url} alt={`${pet.name} ${pet.type} ${index}`} />
            ))
          ) : (
              <div className="no-images-message">
                <p>No images available</p>
                <button className="add-photo">Click here to add Photo</button>
              </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default PetPage;
