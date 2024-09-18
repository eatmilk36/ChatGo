package WebSocket

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

// 向群組中的所有客戶端發送訊息
func (g *Group) Run() {
	for {
		msg := <-g.Broadcast
		for client := range g.Clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(g.Clients, client)
			}
		}
	}
}
