import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import '../css/Login.css';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const backendUrl = process.env.REACT_APP_BACKEND_BASE_URL;
    const navigate = useNavigate();
    console.log(backendUrl);

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleLogin = () => {
        fetch(backendUrl + 'login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error('Login gagal!');
                }
            })
            .then(data => {
                navigate('/dashboard');
            })
            .catch(error => {
                setErrorMessage('Terjadi kesalahan. Silakan coba lagi.');
                console.error(error);
            });
    };

    return (
        <div className="container">
            <h1>Login</h1>
            <div className="form-group">
                <input
                    className="input-field"
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={handleUsernameChange}
                />
            </div>
            <div className="form-group">
                <input
                    className="input-field"
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={handlePasswordChange}
                />
            </div>
            <button className="button" onClick={handleLogin}>Login</button>
            {errorMessage && <div className="error-message">{errorMessage}</div>}
        </div>
    );
};

export default Login;
