package Services

import (
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

type LogELKService struct{}

func NewLogELKService() *LogELKService {
	//conn, err := gelf.NewWriter("localhost:5044")
	conn, err := net.Dial("tcp", "logstash:5044")
	if err != nil {
		logrus.Fatalf("Failed to create gelf writer: %v", err)
	}

	// 設定 Logrus 使用 gelf 傳送日誌
	logrus.SetOutput(conn)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 設定輸出到標準輸出，或你也可以指定寫入文件
	logrus.SetOutput(os.Stdout)

	// 設定日誌級別（例如 Info, Debug, Error）這邊是設定輸出級別
	logrus.SetLevel(logrus.DebugLevel)

	return &LogELKService{}
}

func (Log *LogELKService) LogError(value string) {
	logrus.WithFields(logrus.Fields{
		"event": "error_event",
		"user":  "user1",
	}).Error(value)
}

func (Log *LogELKService) LogDebug(value string) {
	logrus.WithFields(logrus.Fields{
		"event": "debug_event",
		"user":  "user2",
	}).Debug(value)
}

func (Log *LogELKService) LogInfo(value string) {
	logrus.WithFields(logrus.Fields{
		"event": "info_event",
		"user":  "user3",
	}).Info(value)
}
