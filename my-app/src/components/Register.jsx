import React, { useState } from "react";
import "./LoginStyle.css";
import Navbar from "./Navbar";
import { Link } from "react-router-dom";

export const Register = (props) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');
    const [passwordConfirm, setPasswordConfirm] = useState('');
    const [showMessage, setShowMessage] = useState(false); // State variable to manage message visibility

    async function handleSubmit(event) {
        event.preventDefault();
        let item = { name, email, password, passwordConfirm };
        let result = await fetch("http://localhost:8080/v1/signup", {
            method: 'POST',
            body: JSON.stringify(item),
            Headers: {
                "Content-Type": 'application/json',
                "Accept": 'application/json'
            }
        });
        result = await result.json();
        console.warn("result", result);
        setShowMessage(true); // Show message after form submission
    }

    return (
        <div className="register">
            <div className="auth-form-container">
                <Navbar />
                {showMessage ? (
                    <div className="message">
                        <p>Please verify your account (Check your email for further instructions) </p> 
                        <p>Check your email for further instructions or Click <Link to="/login">here</Link> to sign in.</p>
                    </div>
                ) : (
                    <>
                        <h2>Register</h2>
                        <form className="register-form" onSubmit={handleSubmit}>
                            <label htmlFor="name">Name</label>
                            <input value={name} onChange={(e) => setName(e.target.value)} type="name" placeholder="YourName" id="name" name="name" />
                            <label htmlFor="email">Email</label>
                            <input value={email} onChange={(e) => setEmail(e.target.value)} type="email" placeholder="youremail@domain.com" id="email" name="email" />
                            <label htmlFor="password">Password</label>
                            <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="*********" id="password" name="password" />
                            <label htmlFor="passwordConfirm">Confirm Password</label>
                            <input value={passwordConfirm} onChange={(e) => setPasswordConfirm(e.target.value)} type="password" placeholder="*********" id="passwordConfirm" name="passwordConfirm" />
                            <button type="submit">Register</button>
                        </form>
                        <Link to="/login">
                            <button className="link-btn" type="button">Already have an account? Login.</button>
                        </Link>
                    </>
                )}
            </div>
        </div>
    )
}
