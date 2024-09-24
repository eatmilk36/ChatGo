package User

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// NewUserRepository 回傳一個新的 UserRepository 實例
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *Model) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id uint) (*Model, error) {
	var user Model
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *Repository) UpdateUser(user *Model) error {
	return r.db.Save(user).Error
}

func (r *Repository) DeleteUser(id uint) error {
	return r.db.Delete(&Model{}, id).Error
}

func (r *Repository) GetUserByAccountAndPassword(account string, password string) (*Model, error) {
	var user Model
	if err := r.db.Where("account = ? AND password = ?", account, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
