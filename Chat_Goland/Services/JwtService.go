package Services

import (
	"Chat_Goland/Single/SingleConfig"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
	"time"
)

type JwtService struct{}

// MyCustomClaims 定義自訂的 Claim 結構
type MyCustomClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 產生JWT Token
func (j JwtService) GenerateJWT(username string, id int) (string, error) {

	// 設定過期時間
	expirationTime := time.Now().Add(1 * time.Hour)

	// 建立自訂的 Claims
	claims := MyCustomClaims{
		UserId:   strconv.Itoa(id),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "chat",
		},
	}

	// 建立 Token，使用 HS256 演算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 簽名 Token 並返回作為字串
	config := SingleConfig.SingleConfig
	tokenString, err := token.SignedString([]byte(config.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JwtService) ValidateJWT(tokenString string) (*MyCustomClaims, error) {
	// 解析並驗證 JWT
	token, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 確認使用的簽名方法是否正確
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		config := SingleConfig.SingleConfig
		return []byte(config.Jwt.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 檢查 token 是否有效以及是否有自訂 Claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
