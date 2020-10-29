package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
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

func NewRedisEngine(r *setting.RedisSettings) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   r.MaxIdle,
		MaxActive: r.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", r.Host)
			if err != nil {
				return nil, err
			}
			if r.Password != "" {
				if _, err := c.Do("AUTH", r.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}
