import React, { useState } from 'react';

const Login = () => {
    // 定義表單狀態
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    // 處理表單提交
    const handleSubmit = (event) => {
        event.preventDefault(); // 防止頁面刷新

        // 簡單的驗證
        if (username === '' || password === '') {
            setError('請輸入使用者名稱和密碼');
        } else {
            setError('');
            console.log('登入中...', { username, password });
            // 這裡可以發送 API 請求來進行後端驗證
        }
    };

    return (
        <div style={{ maxWidth: '400px', margin: '0 auto' }}>
            <h2>登入</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="username">使用者名稱</label>
                    <input
                        type="text"
                        id="username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        placeholder="輸入使用者名稱"
                    />
                </div>
                <div>
                    <label htmlFor="password">密碼</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="輸入密碼"
                    />
                </div>
                {error && <p style={{ color: 'red' }}>{error}</p>}
                <button type="submit">登入</button>
            </form>
        </div>
    );
};

export default Login;