package ChatroomMessageHistory

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// NewChatroomMessageHistoryRepository 回傳一個新的 ChatroomMessageHistoryRepository 實例
func NewChatroomMessageHistoryRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateChatroomMessageHistoryRepository(chatroomMessageHistoryRepository []Model) error {
	return r.db.Create(chatroomMessageHistoryRepository).Error
}

func (r *Repository) GetChatroomMessageHistoryRepositoryByID(id uint) (*Model, error) {
	var chatroomMessageHistoryRepository Model
	err := r.db.First(&chatroomMessageHistoryRepository, id).Error
	return &chatroomMessageHistoryRepository, err
}

func (r *Repository) UpdateChatroomMessageHistoryRepository(chatroomMessageHistoryRepository *Model) error {
	return r.db.Save(chatroomMessageHistoryRepository).Error
}

func (r *Repository) DeleteChatroomMessageHistoryRepository(id uint) error {
	return r.db.Delete(&Model{}, id).Error
}

func (r *Repository) GetChatroomMessageHistoryRepositoryByAccountAndPassword(account string, password string) (*Model, error) {
	var chatroomMessageHistoryRepository Model
	if err := r.db.Where("account = ? AND password = ?", account, password).First(&chatroomMessageHistoryRepository).Error; err != nil {
		return nil, err
	}
	return &chatroomMessageHistoryRepository, nil
}

func (r *Repository) GetLastTimeStamp() int64 {
	var model Model
	result := r.db.Order("TimeStamp DESC").First(&model)
	if result.Error != nil {
		return 0
	}
	return model.TimeStamp
}
