package WebSocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var groupManager = NewGroupManager()

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允許來自任何來源的請求
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升級失敗:", err)
		return
	}
	defer conn.Close()

	// 假設這裡從URL或訊息中獲取客戶端想要加入的群組
	groupName := r.URL.Query().Get("group")
	if groupName == "" {
		groupName = "default"
	}

	groupManager.JoinGroup(groupName, conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("讀取訊息錯誤:", err)
			break
		}

		// 轉發訊息給群組中的所有人
		groupManager.SendToGroup(groupName, message)
	}
}
