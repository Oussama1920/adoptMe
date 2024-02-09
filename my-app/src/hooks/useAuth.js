// useAuth.js
import { useState } from 'react';

const useAuth = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  // Function to handle login
  const login = (accessToken) => {
    localStorage.setItem('token', accessToken);
    setIsAuthenticated(true);
  };

  // Function to handle logout
  const logout = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
  };

  // Function to check if the user is authenticated
  const checkAuth = () => {
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token); // Set isAuthenticated to true if token exists, false otherwise
  };

  // Call checkAuth when the component mounts to check if the user is already authenticated
  useState(checkAuth);

  return { isAuthenticated, login, logout, checkAuth };
};

export default useAuth;
