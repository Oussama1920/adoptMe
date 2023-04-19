import React, {useState} from "react";
import './App.css';
import {Login} from "./components/Login";
import {Register} from "./components/Register";
import Navbar from "./components/Navbar";
import {Route, Routes} from "react-router-dom";
import Home from "./routes/Home";
import Signup from "./routes/Signup";
import About from "./routes/About";
import Contact from "./routes/Contact";

function App() {
  const [currentForm,setCurrentForm] = useState('login');
  
  const toggleForm = (forName) => {
    setCurrentForm(forName);
  }
// {currentForm === "login" ? <Login  onFormSwitch={toggleForm}/> : <Register onFormSwitch={toggleForm}/>} 
  return (
    <div className="App">
       <Routes>
       <Route path="/" element={<Home/>}/>
        <Route path="/about" element={<About/>}/>
        <Route path="/contact" element={<Contact/>}/>
        <Route path="/signup" element={<Signup/>}/>
       </Routes>
    </div>
  );
}

export default App;
