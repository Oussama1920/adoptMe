// PetSearchForm.jsx
import React, { useState } from 'react';

const PetSearchForm = ({ onSearch }) => {
  const [searchParams, setSearchParams] = useState({
    name: '',
    age: '',
    type: '',
    createdBefore: '',
    createdAfter: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setSearchParams(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSearch(searchParams);
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Name:
        <input type="text" name="name" value={searchParams.name} onChange={handleChange} />
      </label>
      {/* Other input fields for age, type, createdBefore, createdAfter */}
      <button type="submit">Search</button>
    </form>
  );
};

export default PetSearchForm;
