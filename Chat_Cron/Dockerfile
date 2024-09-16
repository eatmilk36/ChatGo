# 使用 Go 的官方映像
FROM golang:1.22

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 複製所有 Go 源代碼
COPY . ./

RUN ls -al /app

# 編譯 Go 程式
RUN go build -o myapp-cron .

# 確保 myapp-cron 擁有執行權限
RUN chmod +x myapp-cron

# 設定執行命令
CMD ["./myapp-cron"]