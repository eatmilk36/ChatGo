import React from 'react';
import ReactDOM from 'react-dom/client'; // 使用新的 createRoot
import './index.css';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App.js';

const root = ReactDOM.createRoot(document.getElementById('root')); // 使用 createRoot
root.render(
    <React.StrictMode>
        <Router>
            <App />
        </Router>
    </React.StrictMode>
);