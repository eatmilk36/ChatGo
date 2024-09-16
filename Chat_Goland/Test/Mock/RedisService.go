package Mock

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type RedisService struct {
	mock.Mock
}

func (r *RedisService) SaveUserLogin(ctx context.Context, username string, jwt string) error {
	args := r.Called(ctx, username, jwt)
	return args.Error(0)
}
