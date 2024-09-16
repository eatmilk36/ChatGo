package Repositories

import (
	"chat/Config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormUserRepository struct {
	db *gorm.DB
}

func (repo GormUserRepository) InitDatabase() *gorm.DB {
	config, err2 := Config.LoadConfig()
	if err2 != nil {
		panic("Config file load failed")
	}
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.MySql.UserName, config.MySql.Password, config.MySql.Address, config.MySql.Port, config.MySql.Database)
	var err error
	repo.db, err = gorm.Open(mysql.Open(addr))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db, err := repo.db.DB()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil
	}

	db.SetConnMaxLifetime(time.Duration(config.MySql.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(config.MySql.MaxIdleConnections)
	db.SetMaxOpenConns(config.MySql.MaxOpenConnections)

	return repo.db
}
