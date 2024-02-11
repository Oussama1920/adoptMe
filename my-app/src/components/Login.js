import { useRef, useState, useEffect } from 'react';
import useAuth from '../hooks/useAuth';
import { Link, useNavigate } from 'react-router-dom';
import Navbar from './Navbar';
import './LoginStyle.css'; // Import CSS file for Login component styling
import { AiOutlineCloseCircle } from 'react-icons/ai'; // Import the error icon from react-icons library

const LOGIN_URL = '/v1/login';

export const Login = () => {
    const navigate = useNavigate();
    const { login } = useAuth(); // Extract the login function from the useAuth hook

    const userRef = useRef();
    const errRef = useRef();

    const [email, setEmail] = useState('');
    const [password, setPwd] = useState('');
    const [errMsg, setErrMsg] = useState('');

    useEffect(() => {
        userRef.current.focus();
    }, [])

    useEffect(() => {
        setErrMsg('');
    }, [email, password])

    const handleSubmit = async (e) => {
        e.preventDefault();

        let item = { email, password };

        try {
            let response = await fetch("http://localhost:8080/v1/login", {
                method: 'POST',
                body: JSON.stringify(item),
                headers: {
                    "Content-Type": 'application/json',
                    "Accept": 'application/json'
                }
            });

            if (!response.ok) {
                if (response.status === 403) {
                    throw new Error('Please verify your email (check your email)');
                } else if (response.status === 400) {
                    throw new Error('Invalid email or password');
                } else if (response.status === 404) {
                    throw new Error('not user registred with this email');    
                } else {
                    throw new Error('Failed to fetch');
                }
            }

            let responseData = await response.json();
            const accessToken = responseData?.token;
            login(accessToken);
            navigate('/');
        } catch (error) {
            setErrMsg(error.message);
        }
    }

    return (
        <section className="login">
            <div className="auth-form-container">
                <Navbar />
                {/* Apply fade-in animation to the error message */}
                {errMsg && (
                <div className="error-message fade-in">
                    <AiOutlineCloseCircle className="error-icon" size={20} />
                    <p aria-live="assertive">
                        {errMsg}
                    </p>
                </div>
            )}
               { /* 
                

                */ }
                <h2>Login</h2>
                <form className="login-form" onSubmit={handleSubmit}>
                    <label htmlFor="email">Email:</label>
                    <input
                        type="text"
                        id="email"
                        ref={userRef}
                        autoComplete="off"
                        onChange={(e) => setEmail(e.target.value)}
                        value={email}
                        required
                    />

                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        onChange={(e) => setPwd(e.target.value)}
                        value={password}
                        required
                    />
                    <button>Sign In</button>
                </form>
                <Link to="/signup">
                    <button className="link-btn" type="button">Don't have an account? Register here.</button>
                </Link>
            </div>
        </section>
    )
}
