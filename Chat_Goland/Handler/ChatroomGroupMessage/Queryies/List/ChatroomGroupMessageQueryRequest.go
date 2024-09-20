package ChatroomGroupMessage

type ChatroomGroupMessageQueryRequest struct {
	GroupName string `form:"groupName" binding:"required"`
}
