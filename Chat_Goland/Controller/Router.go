package Controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func RouterInit(server *gin.Engine) {
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	test := server.Group("/test")
	{
		test.GET("hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(http.StatusOK, gin.H{"hello": name})
		})
	}

	user := server.Group("/user")
	{
		user.POST("/Login", UserController{}.GetUser)
		user.POST("/Create", UserController{}.CreateUser)
	}

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}
