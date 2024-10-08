package Chatroom

import (
	"Chat_Goland/Interface"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type ChatListQueryHandler struct {
	redis Interface.RedisServiceInterface
	log   Interface.LogServiceInterface
}

func NewChatListQueryHandler(redis Interface.RedisServiceInterface, log Interface.LogServiceInterface) *ChatListQueryHandler {
	return &ChatListQueryHandler{redis: redis, log: log}
}

func (h *ChatListQueryHandler) GetChatroomList(c *gin.Context) {
	ctx := context.Background()
	list, err := h.redis.GetChatList(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	marshal, _ := json.Marshal(list)
	h.log.LogDebug("GetChatList:" + string(marshal))

	//jsonData, _ := json.Marshal(list)
	//c.JSON(http.StatusOK, string(jsonData))
	c.JSON(http.StatusOK, list)
}
