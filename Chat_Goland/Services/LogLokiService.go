package Services

import (
	"log"
	"os"
)

// LogLokiService 定義
type LogLokiService struct {
	logger *log.Logger
}

// NewLogLokiService 創建 LogLokiService 並初始化日誌輸出到文件
func NewLogLokiService() *LogLokiService {
	// 打開或創建日誌文件
	file, err := os.OpenFile("/var/log/myapp/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 將日誌輸出到文件
	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	return &LogLokiService{logger: logger}
}

// LogError 記錄 Error 等級的日誌
func (Log *LogLokiService) LogError(value string) {
	Log.logger.Printf("[ERROR] %s", value)
}

// LogDebug 記錄 Debug 等級的日誌
func (Log *LogLokiService) LogDebug(value string) {
	Log.logger.Printf("[DEBUG] %s", value)
}

// LogInfo 記錄 Info 等級的日誌
func (Log *LogLokiService) LogInfo(value string) {
	Log.logger.Printf("[INFO] %s", value)
}
