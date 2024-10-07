package Create

import (
	"Chat_Goland/Interface"
	"Chat_Goland/Repositories/Models/MySQL/User"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	userRepo Interface.UserRepositoryInterface
	crypto   Interface.CryptoServiceInterService
	log      Interface.LogServiceInterface
}

// NewLoginHandler 建立 CreateHandler 並注入 UserRepositoryInterface
func NewLoginHandler(userRepo Interface.UserRepositoryInterface, crypto Interface.CryptoServiceInterService, log Interface.LogServiceInterface) *Handler {
	return &Handler{userRepo: userRepo, crypto: crypto, log: log}
}

func (h *Handler) CreatUserCommand(c *gin.Context) {
	var req UserCreateRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	model := User.Model{
		Account:     req.Account,
		Password:    h.crypto.Md5Hash(req.Password),
		CreatedTime: req.Createdtime,
	}
	err := h.userRepo.CreateUser(&model)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Create Model Failed")
		return
	}

	marshal, _ := json.Marshal(model)
	h.log.LogDebug("Create Model Success:" + string(marshal))

	c.JSON(http.StatusOK, "ok")
}
