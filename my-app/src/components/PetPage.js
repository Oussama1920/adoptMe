import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Navbar from "../components/Navbar"
import "./PetPageStyle.css";
// todo : changing using JSON to fetch image to URLs ( it seems all websites are using this approach)
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
    <Navbar/> 
    <div className="pet-ad">
      <h1>{pet.name}</h1>
      <p>Type: {pet.type}</p>
      <p>Age: {pet.age}</p>
      <p>createdAt: {pet.created_at}</p>      
      <div>
        {pet.images && pet.images.length > 0 ? (
          pet.images.map((image, index) => (
            <img key={index} src={image.data_url} alt={`${pet.name} ${pet.type} ${index}`} />
          ))
        ) : (
          <p>No images available</p>
        )}
      </div>
    </div>
    </div>
  );
};

export default PetPage;
