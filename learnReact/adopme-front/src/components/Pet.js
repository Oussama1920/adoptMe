import React from 'react';
import './PetStyle.css';

function Pet(props) {
    return (
        <li className='pet'>
            <div className={`availability-bubble ${props.available ? 'available' : 'adopted'}`}>
                {props.available ? 'Disponible' : 'Adopted'}
            </div>
            <p className='name'>{props.name}</p>
            <img src={props.photo} alt={props.name} className='pet-img' />
        </li>
    );
}

export default Pet;
