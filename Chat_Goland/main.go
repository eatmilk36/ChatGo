// @title Gin Swagger API

// @version 1.0

// @host localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /
package main

import (
	"chat/Controller"
	"chat/Middleware"
	"chat/Redis"
	_ "chat/docs"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func main() {
	// 啟動一個 goroutine 來監聽過期事件
	redis := Redis.NewRedisService()
	go redis.ListenForExpiredKeys(context.Background())

	// 啟動中間層檢查JWT
	server := gin.Default()
	server.Use(Middleware.JWTAuthMiddleware())

	// 啟動Router
	Controller.RouterInit(server)
}
