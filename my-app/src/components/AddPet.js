import "./AddPetStyle.css"
import React ,{ useState } from 'react';
import Navbar from './Navbar';

import ImageUploading from "react-images-uploading";

const Results =() => {
  const [images, setImages] = React.useState([]);
  const maxNumber = 69; // max number of image can be added

  const [name, setName] = useState('');
  const [type, setType] = useState('');
  const [age, setAge] = useState('');


  const onChange = (imageList, addUpdateIndex) => {
    // data for submit
    console.log(imageList, addUpdateIndex);
    setImages(imageList);
  };
  async function handleSubmit(event) {
    event.preventDefault();
    let item = { name, type, age, images };
    let result = await fetch("http://localhost:8080/v1/pets/pet", {
        method: 'POST',
        body: JSON.stringify(item),
        headers: {
            "Content-Type": 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    });
        
    try {
        const data = await result.json(); // Await parsing the response as JSON
        console.log("result", data); // Log the parsed JSON data
        if (data.status==="success") {
          console.log("Redirecting to the pet page...");
          // Redirect the user to the pet page with the provided ID
          window.location.href = `/pet/${data.id}`;

        } else {
          console.log("Failed to add the pet. Reloading the page...");
          // Reload the current page to allow the user to recreate the pet
          window.location.reload();
  
        }
    } catch (error) {
        console.error("Failed to parse JSON response", error); // Handle JSON parsing errors
    }

    // Continue with your logic here
}

    return (

        <div className="add-pet">
          <Navbar/>

          <form id="addPetForm" onSubmit={handleSubmit}>

            <label for="name">Name:</label>
            <input value={name} onChange={(e) => setName(e.target.value)} type="name" placeholder="Your Pet Name" id="name" name="name" />

            <label htmlFor="type">Type:</label>
              <select id="type" name="type" required value={type} onChange={(e) => setType(e.target.value)}>
                <option value="">Select Type</option>
                <option value="bird">🐦 Bird</option>
                <option value="cat">🐱 Cat</option>
                <option value="dog">🐶 Dog</option>
              </select>
            
            <label for="age">Age:</label>
            <input value={age} onChange={(e) => setAge(e.target.value)} type="age" placeholder="Your Pet Age" id="age" name="age" />

            <label for="photo">Photos:</label>

            <div className="addPhoto">
                  <ImageUploading
                    multiple
                    value={images}
                    onChange={onChange}
                    maxNumber={maxNumber}
                    dataURLKey="data_url"
                  >
                    {({
                      imageList,
                      onImageUpload,
                      onImageRemoveAll,
                      onImageUpdate,
                      onImageRemove,
                      isDragging,
                      dragProps,
                    }) => (
                      // write your building UI
                      <div className="upload__image-wrapper">
                        <button
                        type="button"
                          style={isDragging ? { color: 'red' } : undefined}
                          onClick={onImageUpload}
                          {...dragProps}
                        >
                          Click or Drop here
                        </button>
                        &nbsp;
                        <button type="button" onClick={onImageRemoveAll}>Remove all images</button>
                        {imageList.map((image, index) => (
                          <div key={index} className="image-item">
                            <img src={image['data_url']} alt="" width="100" />
                            <div className="image-item__btn-wrapper">
                              <button type="button" onClick={() => onImageUpdate(index)}>Update</button>
                              <button type="button" onClick={() => onImageRemove(index)}>Remove</button>
                            </div>
                          </div>
                        ))}
                      </div>
                    )}
                  </ImageUploading>
                </div>
                <button className="button-annonce">publish ad</button>

          </form>
        </div>
    )
}
export default Results