import React, {useState} from "react"
import props from 'prop-types';
import Cookies from 'js-cookie';
import "./LoginStyle.css"
import Navbar from "./Navbar";
import { Link } from "react-router-dom";

export const Login =(props) => {
   const [email, setEmail] = useState('')
   const [password, setPass] = useState('')
   async function handleSubmit (event) {
    event.preventDefault();
    console.log(email);
    let item = {email,password}
    let result = await fetch("http://localhost:8080/auth/login", {
        method: 'POST',
        body: JSON.stringify(item),
        credentials: "same-origin",
        Headers:{
            "Content-Type" :'application/json',
            "Accept":'application/json'
        }
    })
    console.warn("header",result.headers)

    result = await result.json()

    console.warn("result",result)
    Cookies.set('token', "", { expires: 7 });

   }
    return (
               
        
        
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
       
    )
}
