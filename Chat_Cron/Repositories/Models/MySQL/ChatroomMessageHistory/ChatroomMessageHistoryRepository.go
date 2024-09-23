package ChatroomMessageHistory

import (
	"gorm.io/gorm"
)

type GormChatroomMessageHistoryRepository struct {
	db *gorm.DB
}

// NewGormChatroomMessageHistoryRepository 回傳一個新的 GormChatroomMessageHistoryRepository 實例
func NewGormChatroomMessageHistoryRepository(db *gorm.DB) *GormChatroomMessageHistoryRepository {
	return &GormChatroomMessageHistoryRepository{db: db}
}

func (r *GormChatroomMessageHistoryRepository) CreateChatroomMessageHistoryRepository(chatroomMessageHistoryRepository []ChatroomMessageHistoryRepository) error {
	return r.db.Create(chatroomMessageHistoryRepository).Error
}

func (r *GormChatroomMessageHistoryRepository) GetChatroomMessageHistoryRepositoryByID(id uint) (*ChatroomMessageHistoryRepository, error) {
	var chatroomMessageHistoryRepository ChatroomMessageHistoryRepository
	err := r.db.First(&chatroomMessageHistoryRepository, id).Error
	return &chatroomMessageHistoryRepository, err
}

func (r *GormChatroomMessageHistoryRepository) UpdateChatroomMessageHistoryRepository(chatroomMessageHistoryRepository *ChatroomMessageHistoryRepository) error {
	return r.db.Save(chatroomMessageHistoryRepository).Error
}

func (r *GormChatroomMessageHistoryRepository) DeleteChatroomMessageHistoryRepository(id uint) error {
	return r.db.Delete(&ChatroomMessageHistoryRepository{}, id).Error
}

func (r *GormChatroomMessageHistoryRepository) GetChatroomMessageHistoryRepositoryByAccountAndPassword(account string, password string) (*ChatroomMessageHistoryRepository, error) {
	var chatroomMessageHistoryRepository ChatroomMessageHistoryRepository
	if err := r.db.Where("account = ? AND password = ?", account, password).First(&chatroomMessageHistoryRepository).Error; err != nil {
		return nil, err
	}
	return &chatroomMessageHistoryRepository, nil
}
