package Mock

import (
	"Chat_Goland/Repositories/Models/RedisModels"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type RedisService struct {
	mock.Mock
}

func (r *RedisService) SetChatList(ctx context.Context, model RedisModels.RedisChatroomModel) error {
	args := r.Called(ctx, model)
	return args.Error(0)
}

func (r *RedisService) GetChatList(ctx context.Context) ([]string, error) {
	args := r.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}

func (r *RedisService) SaveChatMessage(ctx context.Context, groupName, message string) error {
	args := r.Called(ctx, groupName, message)
	return args.Error(0)
}

func (r *RedisService) GetChatMessage(ctx context.Context, groupName string) ([]string, error) {
	args := r.Called(ctx, groupName)
	return args.Get(0).([]string), args.Error(1)
}

func (r *RedisService) SaveUserLogin(ctx context.Context, username string, jwt string) error {
	args := r.Called(ctx, username, jwt)
	return args.Error(0)
}
