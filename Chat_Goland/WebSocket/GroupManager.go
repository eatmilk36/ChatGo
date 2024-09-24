package WebSocket

import "github.com/gorilla/websocket"

type GroupManager struct {
	Groups map[string]*Group
}

func NewGroupManager() *GroupManager {
	return &GroupManager{
		Groups: make(map[string]*Group),
	}
}

// JoinGroup 加入群組
func (gm *GroupManager) JoinGroup(groupName string, client *websocket.Conn) {
	group, exists := gm.Groups[groupName]
	if !exists {
		group = NewGroup(groupName)
		gm.Groups[groupName] = group
		go group.Run() // 開始處理這個群組的訊息
	}
	group.Clients[client] = true
}

// SendToGroup 向群組發送訊息
func (gm *GroupManager) SendToGroup(groupName string, message []byte) {
	group, exists := gm.Groups[groupName]
	if exists {
		group.Broadcast <- message
	}
}
