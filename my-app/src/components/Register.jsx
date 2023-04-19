import React, {useState} from "react"
import props from 'prop-types';
export const Register =(props) => {
   const [email, setEmail] = useState('');
   const [password, setPass] = useState('');
   const [name, setName] = useState(''); 
   const [passwordConfirm, setPassConfirm] = useState('');

   async function handleSubmit (event) {
    event.preventDefault();
    console.log(email);
    let item = {name,email,password,passwordConfirm}
    let result = await fetch("http://localhost:8080/v1/signup", {
        method: 'POST',
        body: JSON.stringify(item),
        Headers:{
            "Content-Type" :'application/json',
            "Accept":'application/json'
        }
    })
    result = await result.json()
    console.warn("result",result)
   }
    return (
        <div className="auth-form-container">
                <h2>Register</h2>
          <form className="register-form" onSubmit={handleSubmit}>
          <label htmlFor ="name">name</label>
            <input value = {name} onChange={(e) => setName(e.target.value)} type="name" placeholder="YourName" id="name" name="name" />
            <label htmlFor ="email">email</label>
            <input value = {email} onChange={(e) => setEmail(e.target.value)} type="email" placeholder="youremail@domain.com" id="email" name="email" />
            <label value = {password} for ="password ">password</label>
            <input value = {password} onChange={(e) => setPass(e.target.value)} type="password" placeholder="*********" id="password" name="password" />
            <label value = {passwordConfirm} for ="passwordConfirm ">passwordConfirm</label>
            <input value = {passwordConfirm} onChange={(e) => setPassConfirm(e.target.value)} type="password" placeholder="*********" id="passwordConfirm" name="passwordConfirm" />

            <button type="submit">Log In</button>
         </form>
         <button className = "link-btn" onClick={()=> props.onFormSwitch('login')}>Already have an account? Login .</button>
       </div>
    )
}