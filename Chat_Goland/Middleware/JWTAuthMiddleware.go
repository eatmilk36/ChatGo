package Middleware

import (
	"Chat_Goland/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 忽略 JWT 驗證
		if c.FullPath() == "/User/Login" ||
			c.FullPath() == "/swagger/*any" ||
			c.FullPath() == "/User/Create" {
			c.Next()
			return
		}

		// 從請求 Header 取得 Authorization Token
		tokenString := c.GetHeader("Authorization")

		// 驗證 Token
		claims, err := Services.Jwt{}.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 將解碼後的 Claims 設定到 Context 中
		c.Set("username", claims)
		c.Next()
	}
}
