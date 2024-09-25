package Services

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type LogService struct{}

func NewLogService() *LogService {
	// 設定 Logrus 的輸出格式為 JSON
	log.SetFormatter(&log.JSONFormatter{})

	// 設定輸出到標準輸出，或你也可以指定寫入文件
	log.SetOutput(os.Stdout)

	// 設定日誌級別（例如 Info, Debug, Error）這邊是設定輸出級別
	log.SetLevel(log.DebugLevel)

	return &LogService{}
}

func (Log *LogService) LogError(value string) {
	log.WithFields(log.Fields{
		"event": "error_event",
		"user":  "user1",
	}).Error(value)
}

func (Log *LogService) LogDebug(value string) {
	log.WithFields(log.Fields{
		"event": "debug_event",
		"user":  "user2",
	}).Debug(value)
}

func (Log *LogService) LogInfo(value string) {
	log.WithFields(log.Fields{
		"event": "info_event",
		"user":  "user3",
	}).Info(value)
}
