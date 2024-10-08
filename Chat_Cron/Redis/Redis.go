package Redis

import (
	"Chat_Cron/Config"
	"Chat_Cron/Repositories/Models/RedisModels"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisService 負責 Redis 操作
type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	config, err := Config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		panic("config loading failed")
	}
	return &RedisService{client: NewRedisClient(&config.Redis)}
}

// NewRedisClient 建立 Redis 客戶端，使用來自 YAML 的配置
func NewRedisClient(config *Config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         config.Address,
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout * time.Second,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
}

// GetValue 獲取 Redis 中的數值
func (r *RedisService) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// SetValue 設定 Redis 中的 key-value 值
func (r *RedisService) SetValue(ctx context.Context, key string, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err() // 0 表示不設置過期時間
	if err != nil {
		return err
	}
	return nil
}

// ListPush 向 List 中推入元素
func (r *RedisService) ListPush(ctx context.Context, key string, values ...string) error {
	err := r.client.LPush(ctx, key, values).Err()
	if err != nil {
		return err
	}
	return nil
}

// ListRange 獲取 List 中的元素
func (r *RedisService) ListRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	result, err := r.client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRemove 刪除 List 中的特定值
func (r *RedisService) ListRemove(ctx context.Context, key string, count int64, value string) error {
	// count 表示要刪除的數量，正數表示從頭部開始刪除，負數表示從尾部開始刪除，0 表示刪除所有匹配的元素
	err := r.client.LRem(ctx, key, count, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// HashSet 設置 Hash 中的字段值
func (r *RedisService) HashSet(ctx context.Context, key, field, value string) error {
	err := r.client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// HashGet 獲取 Hash 中的字段值
func (r *RedisService) HashGet(ctx context.Context, key, field string) (string, error) {
	result, err := r.client.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("field %s does not exist in key %s", field, key)
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func (r *RedisService) SaveChatMessage(ctx context.Context, groupName, message string) error {
	err := r.ListPush(ctx, "chatroom_"+groupName, message)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisService) GetChatMessage(ctx context.Context, groupName string) ([]string, error) {
	listRange, err := r.ListRange(ctx, "chatroom_"+groupName, 0, -1)
	if err != nil {
		return nil, err
	}
	return listRange, nil
}

func (r *RedisService) GetChatList(ctx context.Context) ([]string, error) {
	listRange, err := r.ListRange(ctx, "ChatRoomList", 0, -1)
	if err != nil {
		return nil, err
	}
	return listRange, nil
}

func (r *RedisService) SetChatList(ctx context.Context, model RedisModels.RedisChatroomModel) error {
	jsonData, err := json.Marshal(model)
	if err != nil {
		return err
	}
	err = r.ListPush(ctx, "ChatRoomList", string(jsonData))
	return nil
}

func (r *RedisService) SaveUserLogin(ctx context.Context, username string, jwt string) error {
	// 檢查是否已經存在
	exists, err := r.client.HExists(ctx, "username:"+username, "name").Result()

	if err != nil {
		return err
	}

	if exists {
		// 更新 JWT
		err := r.client.HSet(ctx, "username:"+username, "jwt", jwt).Err()
		if err != nil {
			return err
		}
	} else {
		// 新增用戶和 JWT
		err := r.client.HSet(ctx, "LoginUser:"+username, "username", username, "jwt", jwt).Err()
		if err != nil {
			return err
		}
	}

	// 設置過期時間
	err = r.client.Expire(ctx, "LoginUser:"+username, time.Minute*30).Err()
	if err != nil {
		return err
	}

	return nil
}

// ListenForExpiredKeys 獨立的函數來監聽 Redis 的鍵過期事件
func (r *RedisService) ListenForExpiredKeys(ctx context.Context) {
	PubNub := r.client.PSubscribe(ctx, "__keyevent@0__:expired")
	defer func(PubNub *redis.PubSub) {
		_ = PubNub.Close()
	}(PubNub)

	for msg := range PubNub.Channel() {
		fmt.Println("Key expired:", msg.Payload)
		// 在這裡執行過期鍵的通知邏輯，例如透過 MQTT 或 WebSocket 發送通知
	}
}
