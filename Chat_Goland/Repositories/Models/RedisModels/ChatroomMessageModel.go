package RedisModels

type ChatroomMessageModel struct {
	UserId    int    `json:"userId"`
	GroupName string `json:"groupName"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
