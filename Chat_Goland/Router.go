package main

import (
	"Chat_Goland/Controller"
	"Chat_Goland/Middleware"
	"Chat_Goland/Redis"
	"Chat_Goland/Repositories"
	"Chat_Goland/Repositories/Models/MySQL/User"
	"Chat_Goland/Services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func RouterInit() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:80", "http://172.30.240.1:80", "http://localhost", "http://localhost:80", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	//server.Use(func(c *gin.Context) {
	//	c.Next()
	//	// 檢查 CORS 頭部是否正確被設置
	//	fmt.Println("CORS Origin:", c.Writer.Header().Get("Access-Control-Allow-Origin"))
	//})

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
		userController := InitUserController()
		user.POST("/Login", userController.GetUser)
		user.POST("/Create", userController.CreateUser)
	}

	chatroom := server.Group("/Chatroom")
	{
		chatroomController := InitChatroomController()
		chatroom.GET("/List", chatroomController.GetChatList)
		chatroom.POST("/Create", chatroomController.SetChatList)
		chatroom.GET("/Message", chatroomController.GetGroupMessage)
	}

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}

func InitUserController() *Controller.UserController {
	// 初始化 Repository
	database := Repositories.Repository{}.InitDatabase()
	repository := User.NewUserRepository(database)

	// 初始化 RedisClient
	redis := Redis.NewRedisService()

	// 初始化 CryptoService
	helper := &Services.CryptoService{}

	// 初始化 JwtService
	jwt := &Services.JwtService{}

	// 初始化 LogService
	log := Services.NewLogService()

	return Controller.NewUserController(
		*repository,
		*redis,
		*helper,
		*jwt,
		*log,
	)
}

func InitChatroomController() *Controller.ChatroomController {
	// 初始化 RedisClient
	redis := Redis.NewRedisService()

	return Controller.NewChatroomController(*redis)
}
