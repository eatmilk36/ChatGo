package Login

import (
	"bytes"
	"chat/Repositories/models"
	"chat/Test"
	"chat/Test/Mock"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var db *gorm.DB

func setup(t *testing.T) {
	// 使用共用的 SetupTestDB 和 ResetDB
	db = Test.SetupTestDB(t)
	Test.ResetDB(db)

	// 插入測試資料
	user := models.User{
		Account:     "Jeter",
		Password:    "MD5",
		Id:          1,
		CreatedTime: time.Now(),
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	user = models.User{
		Account:     "Jeter2",
		Password:    "MD5",
		Id:          2,
		CreatedTime: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}
}

func TestUserCreate(t *testing.T) {
	setup(t)

	// 準備測試的請求
	reqBody := LoginRequest{
		Account:  "Jeter",
		Password: "aa",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	// 構建 HTTP 測試請求
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBodyJson))
	//assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// 使用 httptest 構建回應記錄器
	rec := httptest.NewRecorder()

	// Gin Context 構建
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// 模擬使用者資料庫回應
	mockUser := &models.User{
		Account:     "Jeter",
		Password:    "33",
		Id:          1,
		CreatedTime: time.Now(),
	}

	// 模擬 crypto 回應
	crypto := new(Mock.CryptoHelper)
	crypto.On("Md5Hash", "aa").Return("33")

	// 模擬 Db 回應
	userRepo := new(Mock.UserRepository)
	userRepo.On("GetUserByAccountAndPassword", reqBody.Account, "33").Return(mockUser, nil)

	// 模擬 Redis 回應
	redis := new(Mock.RedisService)
	redis.On("SaveUserLogin", mock.Anything, mockUser.Account, "33kk").Return(nil)

	// 模擬 Jwt 回應
	jwt := new(Mock.Jwt)
	jwt.On("GenerateJWT", mockUser.Account).Return("33kk", nil)

	// 創建一個具體的 LoginHandler 實例，並初始化必要的依賴
	handler := NewLoginHandler(userRepo, redis, crypto, jwt)
	handler.LoginQueryHandler(c)

	// 驗證 HTTP 狀態碼
	assert.Equal(t, http.StatusOK, rec.Code)

	// 驗證回應中的 JWT
	var respBody string
	err := json.Unmarshal(rec.Body.Bytes(), &respBody)
	assert.NoError(t, err)

	// 檢查回傳的 JWT 是否與模擬生成的 JWT 一致
	assert.Equal(t, "33kk", respBody)

	// 確認 Mock 的 CryptoHelper, UserRepository, RedisService 和 Jwt 的期望都達成
	crypto.AssertExpectations(t)
	userRepo.AssertExpectations(t)
	redis.AssertExpectations(t)
	jwt.AssertExpectations(t)
}
