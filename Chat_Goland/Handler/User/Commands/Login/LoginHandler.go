package Login

import (
	"chat/Ineterface"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type Handler struct {
	userRepo Ineterface.UserRepository
	redis    Ineterface.RedisServiceInterface
	crypto   Ineterface.CryptoHelper
	jwt      Ineterface.JwtInterface
}

// NewLoginHandler 建立 LoginHandler 並注入 UserRepository
func NewLoginHandler(userRepo Ineterface.UserRepository, redis Ineterface.RedisServiceInterface, crypto Ineterface.CryptoHelper, jwt Ineterface.JwtInterface) *Handler {
	return &Handler{userRepo: userRepo, redis: redis, crypto: crypto, jwt: jwt}
}

func (h *Handler) LoginQueryHandler(c *gin.Context) {
	var req LoginRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 使用解析出的 account 和 password
	user, err := h.userRepo.GetUserByAccountAndPassword(req.Account, h.crypto.Md5Hash(req.Password))

	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	jwt, _ := h.jwt.GenerateJWT(user.Account)

	err = h.redis.SaveUserLogin(context.Background(), user.Account, jwt)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	c.JSON(http.StatusOK, jwt)
}
