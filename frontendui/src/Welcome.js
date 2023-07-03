import React from 'react';
import {Link} from 'react-router-dom';
import './Welcome.css';


function Welcome(){
    return (
        <div className="welcome-container">
            <h1>Welcome to Our Social Media App</h1>
            <p>Connect with friends, share moments, and engage with the community.</p>
            <div className="cta-container">
                <Link to="/login" className="cta-button" >Login</Link>
                <Link to="/signup" className="cta-button">Signup</Link>
            </div>
        </div>
    );
}

export default Welcome;