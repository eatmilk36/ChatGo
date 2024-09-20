import React, {useEffect, useRef, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {getToken} from "../Common/LocalStorage";

function Chatroom() {
    const navigate = useNavigate();
    const [messages, setMessages] = useState([]);
    const socketRef = useRef(null);
    const [inputValue, setInputValue] = useState('');
    const {id} = useParams();

    useEffect(() => {

        let token = getToken();
        if (token === undefined || token === null) {
            navigate('/login');
        }

        // 創建 WebSocket 連接
        socketRef.current = new WebSocket('ws://127.0.0.1:52333/ws?group=' + id);

        // 監聽來自伺服器的訊息
        socketRef.current.onmessage = (event) => {
            const newMessage = event.data;
            setMessages((prevMessages) => [...prevMessages, newMessage]);
        };

        // 清理 WebSocket 連接
        return () => {
            socketRef.current.close();
        };
    }, []);

    // 發送訊息到 WebSocket 伺服器
    const sendMessage = () => {
        console.log(123)
        console.log(socketRef.current)
        console.log(socketRef.current.readyState === WebSocket.OPEN)
        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            console.log('sendMessage')
            socketRef.current.send(inputValue); // 發送輸入框的值
            setInputValue(''); // 清空輸入框
        }
    };
    // 加入群組
    // 需要接收傳過來的參數
    // 傳送與接收Socket資料並顯示
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