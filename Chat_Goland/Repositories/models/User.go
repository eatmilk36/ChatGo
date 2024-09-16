package models

import (
	"time"
)

type User struct {
	//gorm.Model
	Account     string    `json:"Account" gorm:"primaryKey"`
	Password    string    `json:"Password" gorm:"type:varchar(30)"`
	Id          int       `json:"Id" gorm:"type:varchar(30)"`
	CreatedTime time.Time `json:"Createdtime" gorm:"column:Createdtime"`
}
