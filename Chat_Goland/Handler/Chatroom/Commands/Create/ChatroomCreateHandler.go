package Chatroom

import (
	"Chat_Goland/Interface"
	"Chat_Goland/Repositories/Models/RedisModels"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type ChatroomCreateHandler struct {
	redis Interface.RedisServiceInterface
	log   Interface.LogServiceInterface
}

func NewChatroomCreateHandler(redis Interface.RedisServiceInterface, log Interface.LogServiceInterface) *ChatroomCreateHandler {
	return &ChatroomCreateHandler{redis: redis, log: log}
}

func (h *ChatroomCreateHandler) SetChatroomList(c *gin.Context) {
	var req ChatroomCreateHandlerRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	err := h.redis.SetChatList(ctx, RedisModels.RedisChatroomModel{
		Id:   req.Id,
		Hash: req.Hash,
		Name: req.Name,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	h.log.LogDebug("Create Chatroom Success:" + req.Name)

	c.JSON(http.StatusOK, "ok")
}
