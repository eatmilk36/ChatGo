@REM docker-compose down
@REM docker-compose up --build
@REM docker-compose up -d
docker-compose up --build -d

--------------------------
go test
go test -v      //詳細輸出
go test -cover

// 測試詳細報告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

// 使用SQLite
set CGO_ENABLED=1

// swagger 產生文件
swag init

// 如果套件無法匯入用下列指令
go mod init your_module_name
go get github.com/robfig/cron/v3  // 這是要安裝的套件
go mod tidy
// 還是不行就刪除 go.mod 在自己創一個候執行上面的指令

