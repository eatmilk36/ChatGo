package Mock

import (
	"chat/Repositories/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) CreateUser(user *models.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepository) GetUserByID(id uint) (*models.User, error) {
	args := u.Called(id)
	// args.Get(0) 是 *models.User，需要轉換類型
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepository) DeleteUser(id uint) error {
	args := u.Called(id)
	return args.Error(0)
}

func (u *UserRepository) GetUserByAccountAndPassword(account string, password string) (*models.User, error) {
	args := u.Called(account, password)
	user, ok := args.Get(0).(*models.User) // 取得模擬回傳的 *models.User
	if !ok {
		return nil, args.Error(1) // 如果模擬對象不存在，則回傳錯誤
	}
	return user, args.Error(1) // 回傳使用者與錯誤
}
