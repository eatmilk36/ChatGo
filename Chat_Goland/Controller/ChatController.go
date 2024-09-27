package Controller

import (
	ChatroomCreate "Chat_Goland/Handler/Chatroom/Commands/Create"
	ChatroomList "Chat_Goland/Handler/Chatroom/Queries/List"
	ChatroomGroupMessage "Chat_Goland/Handler/ChatroomGroupMessage/Queryies/List"
	"Chat_Goland/Redis"
	"Chat_Goland/Services"
	"github.com/gin-gonic/gin"
)

type ChatroomController struct {
	redisService Redis.RedisService
	log          Services.LogLokiService
}

func NewChatroomController(redis Redis.RedisService, log Services.LogLokiService) *ChatroomController {
	return &ChatroomController{
		redisService: redis,
		log:          log,
	}
}

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
	// 注入到 LoginHandler
	ChatroomList.NewChatListQueryHandler(&ctrl.redisService, &ctrl.log).GetChatroomList(c)
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
	ChatroomCreate.NewChatroomCreateHandler(&ctrl.redisService).SetChatroomList(c)
}

// GetGroupMessage godoc
// @Summary 取得聊天室的群組訊息
// @Description 根據群組名稱取得對應的聊天訊息
// @Tags Chatroom
// @Accept json
// @Produce json
// @Param GroupName query string true "群組名稱"
// @Success 200 {array} []string "成功返回訊息列表"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Router /Chatroom/Message [get]
func (ctrl ChatroomController) GetGroupMessage(c *gin.Context) {
	ChatroomGroupMessage.NewChatroomGroupMessageQueryHandler(&ctrl.redisService).GetChatroomGroupMessage(c)
}
