package Controller

import (
	"Chat_Goland/Middleware"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func RouterInit() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://172.30.240.1:3000", "http://localhost:3000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	server.Use(func(c *gin.Context) {
		c.Next()
		// 檢查 CORS 頭部是否正確被設置
		fmt.Println("CORS Origin:", c.Writer.Header().Get("Access-Control-Allow-Origin"))
	})

	// 啟動中間層檢查JWT
	server.Use(Middleware.JWTAuthMiddleware())

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	test := server.Group("/test")
	{
		test.GET("hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(http.StatusOK, gin.H{"hello": name})
		})
	}

	user := server.Group("/User")
	{
		user.POST("/Login", UserController{}.GetUser)
		user.POST("/Create", UserController{}.CreateUser)
	}

	chatroom := server.Group("/Chatroom")
	{
		chatroom.GET("/List", ChatroomController{}.GetChatList)
		chatroom.POST("/Create", ChatroomController{}.SetChatList)
		chatroom.GET("/Message", ChatroomController{}.GetGroupMessage)
	}

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}
