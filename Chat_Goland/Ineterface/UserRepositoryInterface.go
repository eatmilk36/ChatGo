package Ineterface

import (
	"chat/Repositories/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error

	GetUserByID(id uint) (*models.User, error)

	UpdateUser(user *models.User) error

	DeleteUser(id uint) error

	GetUserByAccountAndPassword(account string, password string) (*models.User, error)
}
