package Ineterface

import (
	"Chat_Goland/Repositories/Models/RedisModels"
	"golang.org/x/net/context"
)

type RedisServiceInterface interface {
	SaveUserLogin(ctx context.Context, username, jwt string) error

	SetChatList(ctx context.Context, model RedisModels.RedisChatroomModel) error

	GetChatList(ctx context.Context) ([]string, error)

	SaveChatMessage(ctx context.Context, groupName, message string) error

	GetChatMessage(ctx context.Context, groupName string) ([]string, error)
}
