package main

import (
	"Chat_Cron/Redis"
	"Chat_Cron/Repositories"
	"Chat_Cron/Repositories/Models/MySQL/ChatroomMessageHistory"
	"Chat_Cron/Repositories/Models/RedisModels"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/context"
	"time"
)

func main() {
	// 創建一個新的 Cron 實例
	c := cron.New()

	// 添加一個定時任務，每分鐘執行一次
	_, _ = c.AddFunc("@every 10s", func() {
		// 之後要備份Redis 到 MySQL
		fmt.Println("每分鐘執行一次的任務：", time.Now())
		fmt.Println("start")

		redis := Redis.NewRedisService()
		ctx := context.Background()
		list, err := redis.GetChatList(ctx)
		if err != nil {
			panic("Redis read chat list fail")
		}

		database := Repositories.GormRepository{}.InitDatabase()

		// 初始化 UserRepository
		chatroomMessageHistoryRepo := ChatroomMessageHistory.NewChatroomMessageHistoryRepository(database)

		// 取得 History message 最後更新時間
		lastTimestamp := chatroomMessageHistoryRepo.GetLastTimeStamp()

		for _, v := range list {
			var chatroom = RedisModels.RedisChatroomModel{}
			err := json.Unmarshal([]byte(v), &chatroom)
			message, err := redis.GetChatMessage(ctx, chatroom.Name)
			if err != nil {
				// log 紀錄失敗
				continue
			}
			// 處存到MySQL
			var histories []ChatroomMessageHistory.Model
			var temp ChatroomMessageHistory.Model
			for _, m := range message {
				err := json.Unmarshal([]byte(m), &temp)
				if err != nil {
					fmt.Println(err.Error() + m)
					continue
				}

				if lastTimestamp < temp.TimeStamp {
					histories = append(histories, temp)
				}
			}

			if len(histories) == 0 {
				continue
			}

			err = chatroomMessageHistoryRepo.CreateChatroomMessageHistoryRepository(histories)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("success")
	})

	// 啟動 Cron 排程
	c.Start()

	// 主程序阻止退出
	select {}
}
