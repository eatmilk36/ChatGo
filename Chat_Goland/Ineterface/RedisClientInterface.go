package Ineterface

import "golang.org/x/net/context"

type RedisServiceInterface interface {
	SaveUserLogin(ctx context.Context, username, jwt string) error
}
