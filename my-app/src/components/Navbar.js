import React, { useState } from 'react';
import "./NavbarStyles.css";
import { MenuItems } from "./Menuitems";
import { Link,useNavigate } from "react-router-dom";
import useAuth from '../hooks/useAuth';
import logoImg from "../assets/logoTadoptini.png"

const Navbar = () => {
    const [clicked, setClicked] = useState(false);
    const { isAuthenticated,logout } = useAuth();
    const navigate = useNavigate();

    const handleClick = () => {
        setClicked(!clicked);
    };
    const handleLogout = () => {
        logout(); // Call the logout function to remove the token and update isAuthenticated
        navigate('/'); // Redirect the user to the home page
    };

    return (
        <nav className="NavbarItems">
            <h1 className="navbar-logo">t'adoptini?</h1>

            <div className="menu-icons" onClick={handleClick}>
                <i className={clicked ? "fas fa-times" : "fas fa-bars"}></i>
            </div>
            <ul className={clicked ? "nav-menu active" : "nav-menu"}>
                {MenuItems.map((item, index) => (
                    <li key={index}>
                        <Link className={item.cName} to={item.url}><i className={item.icon}></i>{item.title}</Link>
                    </li>
                ))}

                {isAuthenticated ? (
                    <div>
                        <button type="button" onClick={handleLogout}>Log out</button>
                        <Link to="/account">
                        <button type="button">my account</button>
                        </Link>
 
                    </div>
                ) : (
                    <div>
                    <Link to="/login">
                        <button type="button">Login</button>
                    </Link>
                     <Link to="/signup">
                        <button type="button">Sign up</button>
                     </Link>
                     </div> 
                )}
                
            </ul>
        </nav>
    );
};

export default Navbar;
