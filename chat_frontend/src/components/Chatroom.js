import React, {useEffect, useRef, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {getToken} from "../Common/LocalStorage.js";
import axios from "../AxiosInterceptors.js";

function Chatroom() {
    const navigate = useNavigate();
    const [messages, setMessages] = useState([]);
    const socketRef = useRef(null);
    const [inputValue, setInputValue] = useState('');
    const {id} = useParams();
    let once = 0;

    useEffect(() => {
        let token = getToken();
        if (!token) {
            navigate('/login');
            return; // 直接返回，避免不必要的 API 呼叫
        }

        // 創建 WebSocket 連接
        if (!socketRef.current) { // 檢查是否已有連接，避免重複連接
            socketRef.current = new WebSocket('ws://127.0.0.1:52333/ws?group=' + id);

            socketRef.current.onmessage = (event) => {
                const newMessage = event.data;
                setMessages((prevMessages) => [...prevMessages, newMessage]);
            };
        }

        if (once === 0) {
            once++;
            axios.get('/Chatroom/Message?groupName=' + id)
                .then((response) => {
                    setMessages((prevMessages) => [...prevMessages, ...response.data]);
                })
                .catch((error) => {
                    console.error("獲取資料時發生錯誤:", error);
                });
        }

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
            socketRef.current.send(inputValue); // 發送輸入框的值
            setInputValue(''); // 清空輸入框
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
            <button onClick={sendMessage}>Send Message</button>
        </div>
    );
}

export default Chatroom;