import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

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
    <div>
      <h1>{pet.name}</h1>
      <p>Type: {pet.type}</p>
      <p>Age: {pet.age}</p>
      <p>createdAt: {pet.created_at}</p>

      <img src={pet.photo} alt={pet.name} />
    </div>
  );
};

export default PetPage;
