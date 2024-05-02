import React from 'react';
import './PetAnnonceStyle.css';
import { Link } from 'react-router-dom';
const getDefaultImage = (type) => {
    switch (type) {
        case 'dog':
            return 'https://www.southernliving.com/thmb/WHH7cdFT3YMJlJN4y7y3lsAKvJ8=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/gettyimages-114166947-1-268128f97e5c415baede328c1fe32f55.jpg';
        case 'cat':
            return 'https://www.thesprucepets.com/thmb/Wy9Vno45XeFtos7omJ80qkZrtZc=/3760x0/filters:no_upscale():strip_icc()/GettyImages-174770333-0f52afc06a024c478fafb1280c1f491f.jpg';
        case 'bird':
            return 'https://rspca.sfo2.cdn.digitaloceanspaces.com/public/Uploads/blog-import/shutterstock_685368169__FocusFillWzEyMDAsNjMwLCJ5IiwxNzNd.jpg';
        default:
            return 'default_image.jpg';
    }
};
function PetAnnonce(props) {
    console.log(" id  is ",props.id)
    return (
        <Link to={`/pet/${props.id}`} className="pet-link">
        <li className='pet'>
            <div className={`availability-bubble ${props.available ? 'available' : 'adopted'}`}>
                {props.available ? 'Disponible' : 'Adopted'}
            </div>
            <p className='name'>{props.name}</p>
            <div>
                {props.images && props.images.length > 0 ? (
                    <ul>
                        {props.images.map((image, index) => (
                            <li key={index}>
                                <img src={image.data_url} alt={`Pet ${index + 1}`} className="pet-img"/>
                            </li>
                        ))}
                    </ul>
                ) : (
                    <img
                        src={getDefaultImage(props.type)}
                        alt={`Default ${props.type} Image`}
                        className="pet-img"
                    />
                )}
            </div>

        </li>
        </Link>

    );
}

export default PetAnnonce;
