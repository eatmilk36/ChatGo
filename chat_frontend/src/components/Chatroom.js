import React, {useEffect, useRef, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {getToken} from "../Common/LocalStorage.js";
import axios from "../AxiosInterceptors.js";
import {jwtDecode} from "jwt-decode";

function Chatroom() {
    const navigate = useNavigate();
    const [messages, setMessages] = useState([]);
    const socketRef = useRef(null);
    const [inputValue, setInputValue] = useState('');
    const {id} = useParams();
    const [isConnected, setIsConnected] = useState(false);

    const connectWebSocket = () => {
        // 創建新的 WebSocket 連接
        // socketRef.current = new WebSocket('ws://127.0.0.1:33925/ws?group=' + id);
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
                            console.log("踢出去")
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

        // socketRef.current.onclose = (event) => {
        //     console.log("WebSocket 連接已關閉，狀態碼:", event.code, "原因:", event.reason);
        //     setIsConnected(false);  // 連接已關閉
        //     // 自動重連
        //     setTimeout(() => {
        //         console.log("嘗試重新連接 WebSocket...");
        //         connectWebSocket();  // 重連
        //     }, 3000);  // 3 秒後重連
        // };

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
        <div>
            <h2>Chatroom</h2>
            <ul>
                {messages.map((msg, index) => (
                    <li key={index}>{msg}</li>
                ))}
            </ul>

            <input
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                placeholder="Enter your message"
            />
            <button onClick={sendMessage} disabled={!isConnected}>Send Message</button>
        </div>
    );
}

export default Chatroom;