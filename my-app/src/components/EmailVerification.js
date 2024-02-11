import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios'; // Import Axios for making HTTP requests
import { AiOutlineLoading } from 'react-icons/ai'; // Import loading spinner from react-icons

export const EmailVerification = () => {
    const { token } = useParams(); // Extract the token from URL params
    const history = useNavigate();
    const [loading, setLoading] = useState(true); // State to track loading status

    useEffect(() => {
        const verifyEmail = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/v1/verifyemail/${token}`);
                const { status, message } = response.data;
                if (status === 'success') {
                    // Email verified successfully
                    console.log(message);
                    // Redirect the user to the sign-in page after a slight delay
                    setTimeout(() => {
                        history('/login');
                    }, 2000); // 2 seconds delay
                } else {
                    // Handle other cases, such as invalid token
                    console.error(message);
                    // Redirect the user to an error page or the home page after a slight delay
                    setTimeout(() => {
                        history('/');
                    }, 2000); // 2 seconds delay
                }
            } catch (error) {
                console.error('Error verifying email:', error);
                // Redirect the user to an error page or the home page after a slight delay
                setTimeout(() => {
                    history('/');
                }, 2000); // 2 seconds delay
            } finally {
                // Set loading to false after verifying email
                setTimeout(() => {
                    setLoading(false);
                }, 2000); // 2 seconds delay
            }
        };

        verifyEmail();
    }, [token, history]);

    return (
        <div>
            {loading ? (
                <div className="loading-spinner">
                    <AiOutlineLoading size={32} className="spinner" />
                    <p>Verifying your email...</p>
                </div>
            ) : (
                <p>Email verification complete!</p>
                // You can display a success message here
            )}
        </div>
    );

};

export default EmailVerification;
