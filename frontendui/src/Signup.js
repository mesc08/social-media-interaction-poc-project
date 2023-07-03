import React, { useState} from 'react';
import { Link } from 'react-router-dom';
import './Login.css';

function Signup() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle API logic

    const isAuthenticated = true;
    if (isAuthenticated) {
      // Redirect to dashboard page
      window.location.href = '/dashboard';
    } else {
      // Redirect to welcome page
      window.location.href = '/welcome';
    }
  };

  return (
    <div className="login-form-container">
      <form className="login-form" onSubmit={handleSubmit}>
        <h2>SignUp</h2>
        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={handleEmailChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={handlePasswordChange}
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Signup
        </button>
        <p>
          Already have an account? <Link to="../login">Login</Link>
        </p>
      </form>
    </div>
  );
}

export default Signup;
