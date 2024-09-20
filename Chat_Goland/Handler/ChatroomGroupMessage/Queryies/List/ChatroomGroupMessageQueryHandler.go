package ChatroomGroupMessage

import (
	"Chat_Goland/Ineterface"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

type ChatroomGroupMessageQueryHandler struct {
	redis Ineterface.RedisServiceInterface
}

func NewChatroomGroupMessageQueryHandler(redis Ineterface.RedisServiceInterface) *ChatroomGroupMessageQueryHandler {
	return &ChatroomGroupMessageQueryHandler{redis: redis}
}

func (h *ChatroomGroupMessageQueryHandler) GetChatroomGroupMessage(c *gin.Context) {
	var req ChatroomGroupMessageQueryRequest

	// 綁定 JSON 參數
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	list, err := h.redis.GetChatMessage(ctx, req.GroupName)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Redis failed")
		return
	}

	c.JSON(http.StatusOK, list)
}
