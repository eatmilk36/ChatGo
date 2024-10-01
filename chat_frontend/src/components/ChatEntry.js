import React, {useEffect, useState} from 'react';
import '../css/ChatEntry.css'; // 加上 CSS 來做簡單的樣式
import axios from "../AxiosInterceptors.js";
import {useNavigate} from 'react-router-dom';
import CryptoJS from 'crypto-js';
import 'bootstrap/dist/css/bootstrap.min.css';

function ChatEntry() {
    const [lists, setLists] = useState([]);
    const [chatroomName, setChatroomName] = useState('');
    const navigate = useNavigate();
    const timestamp = Date.now(); // 取得當前的時間戳 (毫秒)

    // 使用 useEffect 發送 API 請求來獲取資料
    useEffect(() => {
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

    function CreateChatroom() {
        axios.post('/Chatroom/Create', {
            Id: timestamp,
            Hash: CryptoJS.SHA256(chatroomName).toString(CryptoJS.enc.Hex),
            Name: chatroomName
        })
            .then((_) => {
                navigate(0);
            })
            .catch((error) => {
                console.error("創建聊天室時發生錯誤:", error);
                navigate('/login');
            });
    }

    return (
        <div className="container mt-5">
            <div className="card p-4 shadow-sm mb-4">
                <h3 className="text-center mb-4">創建聊天室</h3>
                <div className="input-group mb-3">
                    <input
                        type="text"
                        className="form-control"
                        value={chatroomName}
                        onChange={(e) => setChatroomName(e.target.value)}
                        placeholder="輸入聊天室名稱"
                    />
                    <button
                        className="btn btn-primary"
                        onClick={CreateChatroom}
                    >
                        創建
                    </button>
                </div>
            </div>

            <div className="list-container row">
                {lists.map((list) => (
                    <List id={list.id} name={list.name} key={list.hash}/>
                ))}
            </div>
        </div>
    );
}

function List({id, name}) {
    const navigate = useNavigate();

    const EntryChatroom = () => {
        navigate(`/chat/chatroom/${name}`);
    };

    return (
        <div className="chatroom-list col-md-4 mb-3">
            <div
                className="card p-3 shadow-sm chatroom-item"
                data-key={id}
                onClick={EntryChatroom}
                style={{ cursor: 'pointer' }}
            >
                <h5 className="text-center">{name}</h5>
            </div>
        </div>
    );
}

export default ChatEntry;
