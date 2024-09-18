import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';  // 可以自訂樣式
import Login from './components/Login';  // 引入我們剛剛建立的 Login 元件

// 將 Login 元件掛載到網頁的根元素上
ReactDOM.render(
    <React.StrictMode>
        <Login />
    </React.StrictMode>,
    document.getElementById('root')
);