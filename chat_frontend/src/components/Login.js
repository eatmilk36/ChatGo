import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    // 定義表單狀態
    const [account, setAccount] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    // 處理登入操作
    const handleLogin = async () => {
        // 簡單的驗證
        if (account === '' || password === '') {
            setError('請輸入使用者名稱和密碼');
        } else {
            // setError('');
            // setLoading(true); // 開始加載狀態

            try {
                // 發送 API 請求來進行後端驗證
                const response = await axios.post('http://127.0.0.1:8080/user/Login', {
                    account: account,
                    password: password,
                }, {
                    headers: {
                        'Content-Type': 'application/json', // 明確設定傳送 JSON 資料
                    },
                    withCredentials: true // 如果需要發送 Cookies 或者其他憑證
                });

                if (response.data === "") {
                    alert("登入失敗")
                    return;
                }

                // 驗證成功
                console.log('登入成功:', response.data);
                // 轉頁到Chat
                navigate('/chat');
            } catch (error) {
                // 驗證失敗
                if (error.response) {
                    setError('登入失敗: ' + error.response.data.message);
                } else {
                    setError('登入失敗，請稍後再試');
                }
            } finally {
                setLoading(false); // 結束加載狀態
            }
        }
    };

    function getAccount(e) {
        setAccount(e.target.value);
    }

    function getPassword(e) {
        setPassword(e.target.value);
    }

    return (
        <div style={{ maxWidth: '400px', margin: '0 auto' }}>
            <h2>登入</h2>
            <div>
                <label htmlFor="account">使用者名稱</label>
                <input
                    type="text"
                    id="account"
                    value={account}
                    onChange={(e) => getAccount(e)}
                    placeholder="輸入使用者名稱"
                />
            </div>
            <div>
                <label htmlFor="password">密碼</label>
                <input
                    type="password"
                    id="password"
                    value={password}
                    onChange={(e) => getPassword(e)}
                    placeholder="輸入密碼"
                />
            </div>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <button onClick={handleLogin} disabled={loading}>
                {loading ? '登入中...' : '登入'}
            </button>
        </div>
    );
};

export default Login;