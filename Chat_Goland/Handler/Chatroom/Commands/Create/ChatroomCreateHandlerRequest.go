package Chatroom

type ChatroomCreateHandlerRequest struct {
	Id   int    `json:"id"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}
