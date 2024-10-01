import React, {useEffect, useState} from 'react';
import {Link, Route, Routes, useLocation, useNavigate} from 'react-router-dom';
import Login from './components/Login.js';
import Home from "./components/Home.js";
import ChatEntry from "./components/ChatEntry.js";
import Chatroom from "./components/Chatroom.js";
import {jwtDecode} from "jwt-decode";
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import {getToken, removeToken} from "./Common/LocalStorage.js";

function App() {
    const [isLogin, setIsLogin] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        try {
            let token = getToken();
            if (jwtDecode(token)) {
                setIsLogin(true)
            }
        } catch (error) {
            setIsLogin(false)
            console.log(error)
        }
    }, [location]);

    const Logout = () => {
        removeToken() // 假設你有一個方法來移除 token
        setIsLogin(false);
        navigate('/login');
    };

    return (
        <div className="App">
            <nav className="navbar navbar-expand-lg navbar-light bg-light">
                <div className="container">
                    <Link className="navbar-brand" to="/">Chatroom</Link>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse"
                            data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false"
                            aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav ms-auto">
                            <li className="nav-item">
                                <Link className="nav-link" to="/">Home</Link>
                            </li>
                            <li className="nav-item">
                                <Link className="nav-link" to="/chat/entry">ChatroomEntry</Link>
                            </li>
                            {!isLogin ? (
                                <li className="nav-item">
                                    <Link className="nav-link" to="/login">Login</Link>
                                </li>
                            ) : (
                                <li className="nav-item">
                                    <button className="nav-link btn btn-link" onClick={Logout}>Logout</button>
                                </li>
                            )}
                        </ul>
                    </div>
                </div>
            </nav>

            <div className="container mt-5">
                <Routes>
                    <Route path="/" element={<Home/>}/>
                    <Route path="/login" element={<Login/>}/>
                    <Route path="/chat/entry" element={<ChatEntry/>}/>
                    <Route path="/chat/chatroom/:id" element={<Chatroom/>}/>
                </Routes>
            </div>
        </div>
    );
}

export default App;
