package Interface

import (
	"Chat_Goland/Repositories/Models/MySQL/User"
)

type UserRepositoryInterface interface {
	CreateUser(user *User.Model) error

	GetUserByID(id uint) (*User.Model, error)

	UpdateUser(user *User.Model) error

	DeleteUser(id uint) error

	GetUserByAccountAndPassword(account string, password string) (*User.Model, error)
}
