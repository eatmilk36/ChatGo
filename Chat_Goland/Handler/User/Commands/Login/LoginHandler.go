package Login

import (
	"Chat_Goland/Ineterface"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type Handler struct {
	userRepo Ineterface.UserRepository
	redis    Ineterface.RedisServiceInterface
	crypto   Ineterface.CryptoService
	jwt      Ineterface.JwtServiceInterface
	log      Ineterface.LogServiceInterface
}

func NewLoginHandler(userRepo Ineterface.UserRepository, redis Ineterface.RedisServiceInterface, crypto Ineterface.CryptoService, jwt Ineterface.JwtServiceInterface, log Ineterface.LogServiceInterface) *Handler {
	return &Handler{userRepo: userRepo, redis: redis, crypto: crypto, jwt: jwt, log: log}
}

func (h *Handler) LoginQueryHandler(c *gin.Context) {
	var req LoginRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.LogError("Create User Failed Invalid Request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 使用解析出的 account 和 password
	user, err := h.userRepo.GetUserByAccountAndPassword(req.Account, h.crypto.Md5Hash(req.Password))

	if err != nil || user == nil {
		h.log.LogError("Create User Failed User Not Found")
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	jwt, _ := h.jwt.GenerateJWT(user.Account)

	err = h.redis.SaveUserLogin(context.Background(), user.Account, jwt)

	if err != nil {
		h.log.LogError("Create User Redis Failed")
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	marshal, _ := json.Marshal(user)
	h.log.LogDebug("Create User Success :" + string(marshal))
	c.JSON(http.StatusOK, jwt)
}
