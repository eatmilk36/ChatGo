package Create

import (
	"Chat_Goland/Interface"
	"Chat_Goland/Repositories/Models/MySQL/User"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userRepo Interface.UserRepositoryInterface
	crypto   Interface.CryptoServiceInterService
}

// NewLoginHandler 建立 CreateHandler 並注入 UserRepositoryInterface
func NewLoginHandler(userRepo Interface.UserRepositoryInterface, crypto Interface.CryptoServiceInterService) *Handler {
	return &Handler{userRepo: userRepo, crypto: crypto}
}

func (h *Handler) CreatUserCommand(c *gin.Context) {
	var req UserCreateRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.userRepo.CreateUser(&User.Model{
		Account:     req.Account,
		Password:    h.crypto.Md5Hash(req.Password),
		CreatedTime: req.Createdtime,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, "Create Model Failed")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
