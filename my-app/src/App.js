import React, {Component, useState} from "react";
import './App.css';
import {Login} from "./components/Login";
import {Register} from "./components/Register";
import Navbar from "./components/Navbar";
import {Route, Routes} from "react-router-dom";
import Home from "./routes/Home";
import Signup from "./routes/Signup";
import About from "./routes/About";
import Contact from "./routes/Contact";
import Admin from "./components/Admin";
import RequireAuth from "./components/RequireAuth"

function App() {
  const [currentForm,setCurrentForm] = useState('login');
  const [auth,setAuth] = React.useState(false);
  const toggleForm = (forName) => {
    setCurrentForm(forName);
  }
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
          <Route path="admin" element={<Admin />} />
        </Route>
       </Routes>
    </div>
  );
}
export default App;
