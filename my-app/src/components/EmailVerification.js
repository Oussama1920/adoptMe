import React, { useEffect } from 'react';
import { useParams,  useNavigate } from 'react-router-dom';
import axios from 'axios'; // Import Axios for making HTTP requests

export const EmailVerification = () => {
    const { token } = useParams(); // Extract the token from URL params
    const history = useNavigate();

    useEffect(() => {
        const verifyEmail = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/auth/verifyemail/${token}`);
                const { status, message } = response.data;
                if (status === 'success') {
                    // Email verified successfully
                    console.log(message);
                    // Redirect the user to the sign-in page
                    history('/login');
                } else {
                    // Handle other cases, such as invalid token
                    console.error(message);
                    // Redirect the user to an error page or the home page
                    history('/');
                }
            } catch (error) {
                console.error('Error verifying email:', error);
                // Redirect the user to an error page or the home page
                history('/');
            }
        };

        verifyEmail();
    }, [token, history]);

    return (
        <div>
            <p>Verifying your email...</p>
            {/* You can add a loading spinner or message here if needed */}
        </div>
    );
};

export default EmailVerification;
