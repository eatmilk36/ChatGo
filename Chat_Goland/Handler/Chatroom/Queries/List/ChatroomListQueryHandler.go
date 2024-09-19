package Chatroom

import (
	"Chat_Goland/Ineterface"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type ChatListQueryHandler struct {
	redis Ineterface.RedisServiceInterface
}

func NewChatListQueryHandler(redis Ineterface.RedisServiceInterface) *ChatListQueryHandler {
	return &ChatListQueryHandler{redis: redis}
}

func (h *ChatListQueryHandler) GetChatroomList(c *gin.Context) {
	ctx := context.Background()
	list, err := h.redis.GetChatList(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	//jsonData, _ := json.Marshal(list)
	//c.JSON(http.StatusOK, string(jsonData))
	c.JSON(http.StatusOK, list)
}
