package Services

import (
	"github.com/Graylog2/go-gelf/gelf"
	"log"
	"os"
)

// LogGrayService 定義
type LogGrayService struct {
	writer *gelf.Writer
}

// NewLogGrayService 創建 LogGrayService 並初始化到 Graylog 的連接
func NewLogGrayService() *LogGrayService {
	// 連接到 Graylog 伺服器
	graylogAddr := "graylog:12201" // Graylog 伺服器的地址
	writer, err := gelf.NewWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}

	// 設定預設的 Logger，將 log 發送到 Graylog
	log.SetOutput(writer)

	// 設定標準輸出，確保日誌同時出現在 console
	log.SetOutput(os.Stdout)

	return &LogGrayService{writer: writer}
}

// LogError 記錄 Error 等級的日誌
func (Log *LogGrayService) LogError(value string) {
	msg := &gelf.Message{
		Version:  "1.1",
		Host:     "my-golang-app",
		Short:    value,
		Level:    gelf.LOG_ERR, // Error 等級
		Facility: "my-app",     // 可以設置一個自定義的 facility
	}
	if err := Log.writer.WriteMessage(msg); err != nil {
		log.Printf("Failed to write error log: %v", err)
	}
}

// LogDebug 記錄 Debug 等級的日誌
func (Log *LogGrayService) LogDebug(value string) {
	msg := &gelf.Message{
		Version:  "1.1",
		Host:     "my-golang-app",
		Short:    value,
		Level:    gelf.LOG_DEBUG, // Debug 等級
		Facility: "my-app",
	}
	if err := Log.writer.WriteMessage(msg); err != nil {
		log.Printf("Failed to write debug log: %v", err)
	}
}

// LogInfo 記錄 Info 等級的日誌
func (Log *LogGrayService) LogInfo(value string) {
	msg := &gelf.Message{
		Version:  "1.1",
		Host:     "my-golang-app",
		Short:    value,
		Level:    gelf.LOG_INFO, // Info 等級
		Facility: "my-app",
	}
	if err := Log.writer.WriteMessage(msg); err != nil {
		log.Printf("Failed to write info log: %v", err)
	}
}
