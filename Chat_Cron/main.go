package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 創建一個新的 Cron 實例
	c := cron.New()

	// 添加一個定時任務，每分鐘執行一次
	c.AddFunc("@every 1m", func() {
		// 之後要備份Redis 到 MySQL
		fmt.Println("每分鐘執行一次的任務：", time.Now())
	})

	// 啟動 Cron 排程
	c.Start()

	// 主程序阻止退出
	select {}
}
