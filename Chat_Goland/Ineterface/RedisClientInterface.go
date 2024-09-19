package Ineterface

import (
	"Chat_Goland/Repositories/models/RedisModels"
	"golang.org/x/net/context"
)

type RedisServiceInterface interface {
	SaveUserLogin(ctx context.Context, username, jwt string) error

	SetChatList(ctx context.Context, model RedisModels.RedisChatroomModel) error

	GetChatList(ctx context.Context) ([]string, error)
}
