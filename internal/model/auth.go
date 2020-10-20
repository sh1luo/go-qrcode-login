package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Auth struct {
	AppKey        string     `json:"app_key"`
	AppSecret     string     `json:"app_secret"`
	LastLoginDate *time.Time `json:"last_login_date"`
	LoginIp       string     `json:"login_ip"`
	gorm.Model
}

func (au Auth) Create(db *gorm.DB) error {
	return db.Create(&au).Error
}

func (au Auth) UpdateAccount(db *gorm.DB, updatedMap map[string]interface{}) error {
	return db.Model(au).Update(updatedMap).Error
}

func (au Auth) DelAccount(db *gorm.DB) error {
	return db.Delete(&au).Error
}
