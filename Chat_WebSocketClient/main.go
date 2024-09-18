package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

func main() {
	group1Conn, err := connectToGroup("group1")
	if err != nil {
		log.Fatal("連線失敗:", err)
	}
	defer group1Conn.Close()

	group2Conn, err := connectToGroup("group2")
	if err != nil {
		log.Fatal("連線失敗:", err)
	}
	defer group2Conn.Close()

	// 接收來自 group1 和 group2 的訊息
	go func() {
		for {
			_, message, err := group1Conn.ReadMessage()
			if err != nil {
				log.Println("讀取 group1 訊息錯誤:", err)
				break
			}
			log.Printf("收到group1 訊息: %s", message)
		}
	}()

	go func() {
		for {
			_, message, err := group2Conn.ReadMessage()
			if err != nil {
				log.Println("讀取 group2 訊息錯誤:", err)
				break
			}
			log.Printf("收到group2 訊息: %s", message)
		}
	}()

	go func() {
		for {
			if err := group1Conn.WriteMessage(websocket.TextMessage, []byte("sandy")); err != nil {
				log.Fatal("發送訊息錯誤:", err)
			}

			if err := group1Conn.WriteMessage(websocket.TextMessage, []byte("sara")); err != nil {
				log.Fatal("發送訊息錯誤:", err)
			}

			time.Sleep(10 * time.Second)
		}
	}()

	// 主程序阻止退出
	select {}
}

func connectToGroup(group string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:52333", Path: "/ws", RawQuery: "group=" + group}
	log.Printf("連接到群組 %s: %s", group, u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
