import React  from "react";
import './App.css';
import {Login} from "./components/Login";
import {Register} from "./components/Register";
import {Route, Routes} from "react-router-dom";
import Home from "./routes/Home";
import About from "./routes/About";
import Contact from "./routes/Contact";
import Account from "./components/Account";
import Admin from "./components/Admin";
import RequireAuth from "./components/RequireAuth"
import Pet from "./components/PetPage"
import AddPet from "./components/AddPet";

import {EmailVerification} from "./components/EmailVerification"
function App() {
  
  const ROLES = {
    'User': 2001,
    'Editor': 1984,
    'Admin': 5150
  }
  
// {currentForm === "login" ? <Login  onFormSwitch={toggleForm}/> : <Register onFormSwitch={toggleForm}/>} 
  return (
    <div className="App">
       <Routes>
       <Route path="/" element={<Home/>}/>
        <Route path="/about" element={<About/>}/>
        <Route path="/contact" element={<Contact/>}/>
        <Route path="/signup" element={<Register/>}/>
        <Route path="/login" element={<Login/>}/>
        <Route element={<RequireAuth allowedRoles={[ROLES.Admin]} />}>
        <Route path="admin" element={<Admin/>} /></Route>
        <Route path="/verify-email/:token" element={<EmailVerification/>}/>
        <Route path="/account" element={<Account/>}/>
        <Route path="/pet/:id" element={<Pet/>}/>
        <Route path="/addPet" element={<AddPet/>}/>
       </Routes>
    </div>
  );
}
export default App;
