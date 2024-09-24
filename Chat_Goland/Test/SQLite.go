package Test

import (
	"Chat_Goland/Repositories/Models/MySQL/User"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// SetupTestDB 共用的設置資料庫的函式
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	// 自動遷移
	err = db.AutoMigrate(&User.Model{})
	if err != nil {
		return nil
	}

	return db
}

func ResetDB(db *gorm.DB) {
	err := db.Migrator().DropTable(&User.Model{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&User.Model{})
	if err != nil {
		return
	}
}
