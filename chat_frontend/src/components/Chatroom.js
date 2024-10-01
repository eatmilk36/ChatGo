import React, {useEffect, useRef, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {getToken} from "../Common/LocalStorage.js";
import axios from "../AxiosInterceptors.js";
import {jwtDecode} from "jwt-decode";
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Chatroom.css';

function Chatroom() {
    const navigate = useNavigate();
    const [messages, setMessages] = useState([]);
    const socketRef = useRef(null);
    const [inputValue, setInputValue] = useState('');
    const {id} = useParams();
    const [isConnected, setIsConnected] = useState(false);

    const connectWebSocket = () => {
        // 創建新的 WebSocket 連接
        socketRef.current = new WebSocket('ws://[::1]:33925/ws?group=' + id);

        socketRef.current.onopen = () => {
            console.log("WebSocket 連接已建立");
            setIsConnected(true);
        };

        socketRef.current.onmessage = (event) => {
            const newMessage = event.data;
            if (event.data.startsWith("/AdminKickUser:")) {
                console.log(newMessage);

                let token = getToken();

                if (token) {
                    try {
                        const decodedToken = jwtDecode(token);
                        if (decodedToken.username === event.data.split(':')[2]) {
                            console.log("踢出去");
                            socketRef.current.close();
                            navigate('/login');
                        }

                    } catch (error) {
                        console.error('無效的 JWT:', error);
                    }
                }

                return;
            }
            setMessages((prevMessages) => [...prevMessages, newMessage]);
        };

        socketRef.current.onerror = (error) => {
            console.error("WebSocket 發生錯誤:", error);
        };
    };

    useEffect(() => {
        const token = getToken();
        if (!token) {
            navigate('/login');
            return;
        }

        // 初始化 WebSocket 連接
        connectWebSocket();

        // 獲取初始訊息
        axios.get('/Chatroom/Message?groupName=' + id)
            .then((response) => {
                setMessages((prevMessages) => [...prevMessages, ...response.data]);
            })
            .catch((error) => {
                console.error("獲取資料時發生錯誤:", error);
            });

        // 清理 WebSocket 連接
        return () => {
            if (socketRef.current) {
                socketRef.current.close();
            }
        };
    }, [id, navigate]);

    // 發送訊息到 WebSocket 伺服器
    const sendMessage = () => {
        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            socketRef.current.send(inputValue);  // 發送訊息
            setInputValue('');  // 清空輸入框
        } else {
            console.log("WebSocket 未連接，無法發送訊息");
        }
    };

    return (
        <div className="container mt-5">
            <div className="card chatroom-card shadow-sm">
                <div className="card-header text-center">
                    <h2>聊天室</h2>
                </div>
                <div className="card-body">
                    <ul className="list-group message-list">
                        {messages.map((msg, index) => (
                            <li className="list-group-item" key={index}>
                                {msg}
                            </li>
                        ))}
                    </ul>
                </div>
                <div className="card-footer">
                    <div className="input-group">
                        <input
                            type="text"
                            value={inputValue}
                            onChange={(e) => setInputValue(e.target.value)}
                            className="form-control"
                            placeholder="輸入訊息"
                        />
                        <button
                            onClick={sendMessage}
                            disabled={!isConnected}
                            className="btn btn-primary"
                        >
                            發送訊息
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Chatroom;
