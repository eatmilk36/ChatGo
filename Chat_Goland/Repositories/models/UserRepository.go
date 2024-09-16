package models

import (
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository 回傳一個新的 GormUserRepository 實例
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *GormUserRepository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *GormUserRepository) GetUserByAccountAndPassword(account string, password string) (*User, error) {
	var user User
	if err := r.db.Where("account = ? AND password = ?", account, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
