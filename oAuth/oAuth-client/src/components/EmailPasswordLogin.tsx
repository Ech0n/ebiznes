import React, { useState } from 'react';
import { login, register } from '../api/auth';

function EmailPasswordLogin({ setCurrentPage }:any) {
  const [isRegistering, setIsRegistering] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await login(email, password);
      alert('Login successful!');
      console.log(response);
      if (response && response.token) {
        sessionStorage.setItem('token', response.token);
      }
      setCurrentPage("home")
    } catch (error) {
      alert('Login failed. Please try again.');
    }
  };

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      alert('Passwords do not match.');
      return;
    }
    try {
      const response = await register(email, password);
      alert('Registration successful!');
      console.log(response);
      setIsRegistering(false);
    } catch (error) {
      alert('Registration failed. Please try again.');
    }
  };

  if (isRegistering) {
    return (
      <div>
        <h2>Register</h2>
        <form onSubmit={handleRegister}>
          <div style={styles.fieldContainer}>
            <label>Email:</label>
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          </div>
          <div style={styles.fieldContainer}>
            <label>Password:</label>
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          </div>
          <div style={styles.fieldContainer}>
            <label>Confirm Password:</label>
            <input type="password" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required />
          </div>
          <button type="submit">Register</button>
          <button type="button" onClick={() => setIsRegistering(false)}>Back to Login</button>
        </form>
      </div>
    );
  }

  return (
    <div>
      <h2>Email & Password Login</h2>
      <form onSubmit={handleLogin}>
        <div style={styles.fieldContainer}>
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div style={styles.fieldContainer}>
          <label>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit">Login</button>
      </form>
      <button onClick={() => setIsRegistering(true)}>Create New Account</button>
    </div>
  );
};

const styles = {
  fieldContainer: {
    display: 'flex',
    flexDirection: 'column' as 'column',
  },
};

export default EmailPasswordLogin;
