package Repositories

import (
	"Chat_Goland/Single/SingleConfig"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func (repo Repository) InitDatabase() *gorm.DB {
	config := SingleConfig.SingleConfig

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.MySql.UserName, config.MySql.Password, config.MySql.Address, config.MySql.Port, config.MySql.Database)
	var err error
	maxRetries := 10
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		repo.db, err = gorm.Open(mysql.Open(addr))
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		db, err := repo.db.DB()

		if err == nil {
			db.SetConnMaxLifetime(time.Duration(config.MySql.MaxLifetime) * time.Second)
			db.SetMaxIdleConns(config.MySql.MaxIdleConnections)
			db.SetMaxOpenConns(config.MySql.MaxOpenConnections)
			return repo.db
		}

		time.Sleep(retryDelay)
	}

	return nil
}
