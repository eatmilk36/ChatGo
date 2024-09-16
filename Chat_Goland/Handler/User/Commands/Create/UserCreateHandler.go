package Create

import (
	"chat/Ineterface"
	"chat/Repositories/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userRepo Ineterface.UserRepository
	crypto   Ineterface.CryptoHelper
}

// NewLoginHandler 建立 CreateHandler 並注入 UserRepository
func NewLoginHandler(userRepo Ineterface.UserRepository, crypto Ineterface.CryptoHelper) *Handler {
	return &Handler{userRepo: userRepo, crypto: crypto}
}

func (h *Handler) CreatUserCommand(c *gin.Context) {
	var req UserCreateRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.userRepo.CreateUser(&models.User{
		Account:     req.Account,
		Password:    h.crypto.Md5Hash(req.Password),
		CreatedTime: req.Createdtime,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, "Create User Failed")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
