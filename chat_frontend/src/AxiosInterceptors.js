import axios from 'axios';
import {getToken} from "./Common/LocalStorage.js";

// 創建 Axios 實例
const axiosInstance = axios.create({
    baseURL: 'http://127.0.0.1:8080', // 基本 URL，可以替換成你的 API URL
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: true // 如果需要發送 Cookies 或者其他憑證
});

// 設置請求攔截器
axiosInstance.interceptors.request.use(
    config => {
        const token = getToken(); // 從 localStorage 取得 JWT
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`; // 設定 Authorization header
        }
        return config;
    },
    error => {
        // 請求錯誤時的處理
        return Promise.reject(error);
    }
);

// 設置回應攔截器
axiosInstance.interceptors.response.use(
    response => {
        // 對回應數據做點處理
        return response;
    },
    error => {
        // 統一處理錯誤
        if (error.response && error.response.status === 401) {
            // 未授權錯誤，可能需要重新導向到登入頁面
            console.log('Unauthorized, redirecting to login...');
        }
        return Promise.reject(error);
    }
);

export default axiosInstance;