package GroupManager

import "github.com/gorilla/websocket"

type Group struct {
	Name      string
	Clients   map[*websocket.Conn]bool
	Broadcast chan []byte
}

func NewGroup(name string) *Group {
	return &Group{
		Name:      name,
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
	}
}

// Run 向群組中的所有客戶端發送訊息
func (g *Group) Run() {
	for {
		select {
		case message := <-g.Broadcast:
			// 將訊息發送給群組內的每個客戶端
			for client := range g.Clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					_ = client.Close()
					delete(g.Clients, client) // 移除連線失敗的客戶端
				}
			}
		}
	}
}
