// @title Gin Swagger API

// @version 1.0

// @host localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /
package main

import (
	"Chat_Goland/Controller"
	"Chat_Goland/Middleware"
	"Chat_Goland/Redis"
	"Chat_Goland/WebSocket"
	_ "Chat_Goland/docs"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func main() {
	// 啟動一個 goroutine 來監聽過期事件
	redis := Redis.NewRedisService()
	go redis.ListenForExpiredKeys(context.Background())

	// 啟動中間層檢查JWT
	server := gin.Default()
	server.Use(Middleware.JWTAuthMiddleware())

	// 啟動Router
	go Controller.RouterInit(server)

	// 註冊 WebSocket 處理器
	http.HandleFunc("/ws", WebSocket.WsHandler)

	// 啟動 HTTP 伺服器
	log.Println("伺服器啟動中，監聽端口 :52333")
	err := http.ListenAndServe("127.0.0.1:52333", nil)
	if err != nil {
		log.Fatal("伺服器啟動失敗:", err)
	}

	// 保持運行
	select {}
}
