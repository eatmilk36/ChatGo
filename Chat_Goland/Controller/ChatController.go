package Controller

import (
	ChatroomCreate "Chat_Goland/Handler/Chatroom/Commands/Create"
	ChatroomList "Chat_Goland/Handler/Chatroom/Queries/List"
	"Chat_Goland/Redis"
	"github.com/gin-gonic/gin"
)

type ChatroomController struct{}

// GetChatList godoc
// @Summary Get Chatroom room list
// @Get Chatroom room list
// @Tags Chatroom
// @Accept  json
// @Produce  json
// @Success 200 {object} string "Successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /Chatroom/List [Get]
func (ctrl ChatroomController) GetChatList(c *gin.Context) {

	// 初始化 RedisClient
	redis := Redis.NewRedisService()

	// 注入到 LoginHandler
	handler := ChatroomList.NewChatListQueryHandler(redis)

	// 呼叫 業務邏輯
	handler.GetChatroomList(c)
}

// SetChatList godoc
// @Summary Set Chatroom room list
// @Description Set Chatroom room
// @Tags Chatroom
// @Accept  json
// @Produce  json
// @Param ChatroomCreateHandlerRequest body Chatroom.ChatroomCreateHandlerRequest true "Chatroom credentials"
// @Success 200 {object} string "Successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /Chatroom/Create [Post]
func (ctrl ChatroomController) SetChatList(c *gin.Context) {

	// 初始化 RedisClient
	redis := Redis.NewRedisService()

	// 注入到 LoginHandler
	handler := ChatroomCreate.NewChatroomCreateHandler(redis)

	// 呼叫 業務邏輯
	handler.SetChatroomList(c)
}
