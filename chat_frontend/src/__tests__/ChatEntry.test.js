import React from 'react';
import {render, waitFor} from '@testing-library/react';
import ChatEntry from '../components/ChatEntry';
import axios from "axios";
import {MemoryRouter} from "react-router-dom";
import '@testing-library/jest-dom';


// 模擬 axios
// 在測試中模擬 axios 並避免攔截器的影響
jest.mock('axios', () => {
    return {
        interceptors: {
            request: {
                use: jest.fn(), // 模擬攔截器請求
            },
            response: {
                use: jest.fn(), // 模擬攔截器回應
            }
        },
        get: jest.fn(),
        post: jest.fn(),
        create: jest.fn(function () {
            return this;
        })
    };
});

test('renders user list from API', async () => {
    // 模擬 API 回應
    const mockResponse = {
        data: [
            "{\"id\":1,\"hash\":\"zss\",\"name\":\"jeter\"}",
            "{\"id\":1,\"hash\":\"a1b2c3d4e5\",\"name\":\"Chatroom A\"}",
            "{\"id\":2,\"hash\":\"f6g7h8i9j0\",\"name\":\"Chatroom B\"}",
            "{\"id\":3,\"hash\":\"k1l2m3n4o5\",\"name\":\"Chatroom C\"}"
        ],
    };

    // 模擬 axios 攔截器行為
    axios.interceptors.request.use.mockImplementation((callback) => {
        callback({ headers: { Authorization: 'Bearer token' } });
    });

    axios.interceptors.response.use.mockImplementation(
        (responseCallback, errorCallback) => {
            responseCallback({ data: 'mocked response' });
            errorCallback({ response: { status: 401 } });
        }
    );

    axios.get.mockResolvedValue(mockResponse);

    // 渲染組件
    const {container} = render(
        <MemoryRouter>
            <ChatEntry/>
        </MemoryRouter>
    );

    // 確保組件渲染了來自 API 的資料
    await waitFor(() => {
        const chatroomElement1 = container.querySelector(`[data-key="1"]`);
        // 確保找到該元素
        expect(chatroomElement1).toBeInTheDocument();
        // 從 chatroomElement 中查找 <h3> 元素
        const h3Element = chatroomElement1.querySelector('h3');
        // 確認 <h3> 元素是否存在
        expect(h3Element).toBeInTheDocument();
        // 驗證 <h3> 內的文本是否正確
        expect(h3Element).toHaveTextContent('jeter');
    });
});