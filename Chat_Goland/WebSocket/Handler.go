package WebSocket

import (
	"Chat_Goland/Single/SingleGroupManager"
	"Chat_Goland/Single/SingleRedisServer"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允許來自任何來源的請求
		return true
	},
}

// WsHandler WebSocket 連接處理
func WsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("進入 WebSocket handler")
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket 升級失敗:", err)
		http.Error(w, "WebSocket 升級失敗", http.StatusInternalServerError)
		return
	}
	defer func() {
		log.Println("WebSocket 連接關閉")
		_ = conn.Close()
	}()

	// 假設這裡從URL或訊息中獲取客戶端想要加入的群組
	groupName := r.URL.Query().Get("group")
	if groupName == "" {
		groupName = "default"
	}
	log.Printf("客戶端加入群組: %s", groupName)

	SingleGroupManager.SingleGroupManager.JoinGroup(groupName, conn)
	log.Printf("群組 %s 已加入", groupName)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 檢查 WebSocket 是否異常關閉
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			} else {
				log.Printf("讀取訊息錯誤: %v", err)
			}
			break
		}

		log.Printf("收到訊息: %s", string(message))

		// 儲存訊息到 Redis，使用 goroutine 進行
		go func() {
			if err := saveMessage(groupName, message); err != nil {
				log.Printf("儲存訊息失敗: %v", err)
			} else {
				log.Println("訊息已成功儲存")
			}
		}()

		// 轉發訊息給群組中的所有人
		go SingleGroupManager.SingleGroupManager.SendToGroup(groupName, message)
		log.Printf("訊息已轉發至群組 %s", groupName)
	}
}

// 儲存訊息到 Redis
func saveMessage(groupName string, message []byte) error {
	log.Println("正在儲存訊息到 Redis...")
	ctx := context.Background()

	// 使用 NewRedisService 來初始化 RedisService
	service := SingleRedisServer.SingleRedisServer
	err := service.SaveChatMessage(ctx, groupName, string(message))
	if err != nil {
		log.Printf("儲存訊息至 Redis 失敗: %v", err)
		return err
	}
	log.Println("訊息成功儲存至 Redis")
	return nil
}
