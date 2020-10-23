package model

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sh1luo/go-qrcode-login.git/pkg/setting"
)

func NewDBEngine(dbSetting *setting.MySQLSettings) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			dbSetting.Username,
			dbSetting.Password,
			dbSetting.Host,
			dbSetting.DBName,
			dbSetting.Charset,
			dbSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	db.AutoMigrate(&Auth{}) // 自动迁移表结构
	return db, nil
}

func NewRedisEngine(r *setting.RedisSettings) *redis.Client {
	if rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host,
		Password: r.Password,
		DB:       0,
	}); rdb != nil {
		return nil
	} else {
		return rdb
	}
}
