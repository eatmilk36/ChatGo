import React, {useEffect, useState} from 'react';
import '../css/ChatEntry.css'; // 加上 CSS 來做簡單的樣式
import axios from "../AxiosInterceptors";
import {useNavigate} from 'react-router-dom';

function ChatEntry() {
    const [lists, setLists] = useState([]);
    const navigate = useNavigate();

    // 使用 useEffect 發送 API 請求來獲取資料
    useEffect(() => {
        // 假設你的 API 路徑是 'https://api.example.com/lists'
        axios.get('/Chatroom/List')
            .then((response) => {
                const parsedData = response.data.map(item => JSON.parse(item));
                setLists(parsedData);    // 將解析後的資料保存到狀態中
            })
            .catch((error) => {
                console.error("獲取資料時發生錯誤:", error);
                navigate('/login');
            });
    }, [navigate]);

    return (
        <div className="list-container">
            {lists.map((list) => (
                <List id={list.id} name={list.name} key={list.hash}/>
            ))}
        </div>
    );
}

function List({id, name}) {
    const navigate = useNavigate();

    const EntryChatroom = () => {
        navigate(`/chat/chatroom/${id}`);
    };
    return (
        <div className="chatroom-list" data-key={id} onClick={EntryChatroom}>
            <h3>{name}</h3>
        </div>
    );
}

export default ChatEntry;