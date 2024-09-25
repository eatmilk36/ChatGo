package User

import (
	"time"
)

type Model struct {
	//gorm.Model
	Account     string    `json:"Account" gorm:"primaryKey"`
	Password    string    `json:"Password" gorm:"type:varchar(30)"`
	Id          int       `json:"Id" gorm:"type:varchar(30)"`
	CreatedTime time.Time `json:"Createdtime" gorm:"column:Createdtime"`
}

// TableName 指定資料表名稱為 'ChatroomMessageHistory'
func (Model) TableName() string {
	return "Users"
}
