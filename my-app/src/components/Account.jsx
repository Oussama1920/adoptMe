import { useState, useEffect } from 'react';
import Navbar from "../components/Navbar"
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

function Account() {
    const [user, setUser] = useState({
        email: '',
        name: '',
        firstname: '',
        address: '',
        phonenumber: '',
        dateOfBirth: '',
        created_at: '',
        updated_at: ''
    });

    useEffect(() => {
        // Fetch user information from the backend when the component mounts
        fetchUserData();
    }, []);

    const fetchUserData = async () => {
        try {
            // Make a request to fetch user data from the backend
            const response = await fetch('http://localhost:8080/auth/users/me', {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}` // Include the user's token for authentication
                }
            });
            if (!response.ok) {
                throw new Error('Failed to fetch user data');
            }

            // Parse the response JSON
            const responseData = await response.json();
            const userData = responseData.user; // Access the user object
            console.log("user found is: ",userData)
            // Update the user state with the retrieved data
            setUser(userData);
        } catch (error) {
            console.error('Error fetching user data:', error);
        }
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            // Make a request to update user data on the backend
            const response = await fetch('http://localhost:8080/auth/users/me', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify(user)
            });

            if (!response.ok) {
                throw new Error('Failed to update user data');
            }

            // Show a success message to the user
            alert('User data updated successfully');
        } catch (error) {
            console.error('Error updating user data:', error);
            // Show an error message to the user
            alert('Failed to update user data. Please try again.');
        }
    };

    const handleChange = (event) => {
        const { name, value } = event.target;
        // Update the user state with the new value
        setUser({ ...user, [name]: value });
    };

    return (
        <>
            <section className="login">

            <div className="auth-form-container">
            <Navbar />
                <h2>User Information</h2>
                <form onSubmit={handleSubmit}>
                    <div className="field-group">
                        <label>Email:</label>
                        <input type="text" name="email" value={user.email} readOnly />

                        <label>Name:</label>
                        <input type="text" name="name" value={user.name} onChange={handleChange} />

                        <label>First Name:</label>
                        <input type="text" name="firstname" value={user.firstname} onChange={handleChange} />

                        <label>Address:</label>
                        <input type="text" name="address" value={user.address} onChange={handleChange} />
                    </div>

                    <div className="field-group">
                        <label>Phone Number:</label>
                        <input type="text" name="phonenumber" value={user.phonenumber} onChange={handleChange} />
                            
                        <label>Date of Birth:</label>
                        <DatePicker
                            selected={user.dateOfBirth ? new Date(user.dateOfBirth) : null}
                            onChange={(date) => handleChange({ target: { name: "dateOfBirth", value: date } })}
                            dateFormat="dd/MM/yyyy"
                            isClearable
                        />
                        
                        <label>Created At:</label>
                        <input type="text" name="createdAt" value={user.created_at} readOnly />

                        <label>Updated At:</label>
                        <input type="text" name="updatedAt" value={user.updated_at} readOnly />
                    </div>

                    <button type="submit">Save Changes</button>
                </form>
            </div>
            </section>
        </>
    );
}

export default Account;
