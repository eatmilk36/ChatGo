# 使用 Node.js 官方的基礎映像來構建 React 應用
FROM node:16-alpine as build

# 設置工作目錄
WORKDIR /app

# 複製 package.json 和 package-lock.json
COPY package*.json ./

# 安裝依賴
RUN npm install

# 複製所有檔案
COPY . .

# 構建 React 應用
RUN npm run build

# 使用 Nginx 來提供靜態文件
FROM nginx:alpine

# 複製 React 構建好的靜態文件到 Nginx 的預設公開目錄
COPY --from=build /app/build /usr/share/nginx/html

# 複製自訂的 Nginx 配置(為解決redresh404問題)
COPY ./nginx.conf /etc/nginx/nginx.conf

# 刪除默認的 default.conf 文件，避免衝突(為解決redresh404問題)
RUN rm /etc/nginx/conf.d/default.conf

# Expose the Nginx port
EXPOSE 80

# 啟動 Nginx
CMD ["nginx", "-g", "daemon off;"]