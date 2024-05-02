import React from 'react';
import './PetListStyle.css';
import Pet from './PetAnnonce';

function PetList({ pets }) {
    return (
        <ul className="pets">
            {pets.map((pet, index) => (
                <Pet
                    key={index} // Ensure each component has a unique key
                    type={pet.type}
                    name={pet.name}
                    age={pet.age}
                    available={pet.available}
                    images={pet.images}
                    id={pet.id}

                />
            ))}
        </ul>
    );
}

export default PetList;
