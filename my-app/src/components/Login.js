import { useRef, useState, useEffect } from 'react';
import useAuth from '../hooks/useAuth';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import Navbar from './Navbar';
import axios from '../api/axios';
const LOGIN_URL = '/v1/login';

export const Login = () => {
    const { setAuth } = useAuth();
    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.from?.pathname || "/";

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

        
            let item = {email,password}
            
            let response = await fetch("http://localhost:8080/v1/login", {
            method: 'POST',
            body: JSON.stringify(item),
            Headers:{
            "Content-Type" :'application/json',
            "Accept":'application/json'
        }
        })
            response = await response.json()
            console.warn("result",response)
        
            console.log("HELLO :", JSON.stringify(response));
            console.log(JSON.stringify(response?.token));
            const accessToken = response?.token;
            const roles = "102";
            console.log(roles);
            setAuth({ email, password, roles, accessToken })
            setEmail('');
            setPwd('');
            navigate(from, { replace: true });
        
    }

    return (

        <section className="login">
            <div className="auth-form-container">
            <Navbar/>    
            <p ref={errRef} className={errMsg ? "errmsg" : "offscreen"} aria-live="assertive">{errMsg}</p>
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
                 <button className = "link-btn" type="button">Don't have an account? Register here.</button>
            </Link>
            </div>    
        </section>

    )
}


/*
<div className="login">

<div className="auth-form-container">
<Navbar/>    
 <h2>Login</h2>
  <form className="login-form" onSubmit={handleSubmit}>
    <label htmlFor ="email">email</label>
    <input value = {email} onChange={(e) => setEmail(e.target.value)} type="email" placeholder="youremail@domain.com" id="email" name="email" />
    <label value = {password} for ="password ">password</label>
    <input value = {password} onChange={(e) => setPass(e.target.value)} type="password" placeholder="*********" id="password" name="password" />
    <button type="submit">Log In</button>
 </form>
<Link to="/signup">
    <button className = "link-btn" type="button">Don't have an account? Register here.</button>
</Link>
</div>
</div>

*/
























