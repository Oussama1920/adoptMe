import React, { useState } from 'react';
import "./NavbarStyles.css";
import { MenuItems } from "./Menuitems";
import { Link,useNavigate } from "react-router-dom";
import useAuth from '../hooks/useAuth';

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
                        <button type="button" onClick={handleLogout}>Log out</button> 
                ) : (
                    <Link to="/login">
                        <button type="button">Login</button>
                    </Link>
                )}
            </ul>
        </nav>
    );
};

export default Navbar;
