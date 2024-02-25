import React, { useState } from 'react';
import './searchPetsStyle.css'; // Import the CSS file

const PetSearchAndResults = () => {
  const [searchParams, setSearchParams] = useState({
    name: '',
    age: '',
    type: '',
    created_before: '',
    created_after: ''
  });
  const [searchResults, setSearchResults] = useState([]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setSearchParams(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
  
    // Clear previous search results
    setSearchResults([]);
  
    // Perform search using API
    fetch('http://localhost:8080/v1/pets/pet/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(searchParams)
    })
    .then(response => response.json())
    .then(data => {
      // Set search results
      console.log("result is:",data)
  
      console.log("pets are:",data.pets)
      setSearchResults(data.pets);
    })
    .catch(error => {
      console.error('Error fetching search results:', error);
    });
  };
  

  return (
    <div className="container">
      <form onSubmit={handleSubmit}>
                    <label className="label-container">
                    <span className="label-text">Name:</span>
                 <div className="input-container">
                  <input type="text" name="name" value={searchParams.name} onChange={handleChange} />
                  {searchParams.name && (
                    <button onClick={() => setSearchParams(prevState => ({ ...prevState, name: '' }))}>X</button>
                  )}
                </div>
              </label>
              <label className="label-container">
              <span className="label-text">Age:</span>
                 <div className="input-container">
                  <input type="text" name="age" value={searchParams.age} onChange={handleChange} />
                  {searchParams.age && (
                    <button onClick={() => setSearchParams(prevState => ({ ...prevState, age: '' }))}>X</button>
                  )}
                </div>
              </label>
              <label className="label-container">
              <span className="label-text">Type:</span>
                 <div className="input-container">
                  <input type="text" name="type" value={searchParams.type} onChange={handleChange} />
                  {searchParams.type && (
                    <button onClick={() => setSearchParams(prevState => ({ ...prevState, type: '' }))}>X</button>
                  )}
                </div>
              </label>
              <label className="label-container">
              <span className="label-text">Created Before:</span>
                 <div className="input-container">
                  <input type="date" name="created_before" value={searchParams.created_before} onChange={handleChange} />
                  {searchParams.created_before && (
                    <button onClick={() => setSearchParams(prevState => ({ ...prevState, created_before: '' }))}>X</button>
                  )}
                </div>
              </label>
              <label className="label-container">
              <span className="label-text">Created After:</span>
                 <div className="input-container">
                  <input type="date" name="created_after" value={searchParams.created_after} onChange={handleChange} />
                  {searchParams.created_after && (
                    <button onClick={() => setSearchParams(prevState => ({ ...prevState, created_after: '' }))}>X</button>
                  )}
                </div>
              </label>

        <button type="submit">Search</button>
      </form>

      {/* Display search results */}
      <div>
        <h2>Search Results</h2>
        <ul>
          {searchResults.map(pet => (
            <li key={pet.id}>
              <h3>{pet.name}</h3>
              <p>Age: {pet.age}</p>
              <p>Type: {pet.type}</p>
              <p>Created At: {pet.created_at}</p>
              <div>
                {pet.images && pet.images.length > 0 && (
                  <ul>
                    {pet.images.map((image, index) => (
                      <li key={index}>
                        <img src={image.data_url} alt={`Pet Image ${index + 1}`} />
                      </li>
                    ))}
                  </ul>
                )}
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default PetSearchAndResults;