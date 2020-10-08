package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Auth struct {
	AppKey        string     `json:"app_key"`
	AppSecret     string     `json:"app_secret"`
	Salt          string     `json:"salt"`
	LastLoginDate *time.Time `json:"last_login_date"`
	LoginIp       string     `json:"login_ip"`
	gorm.Model
}

func (au Auth) Create(db *gorm.DB) error { return db.Create(&au).Error }

func (au Auth) UpdateAccount(db *gorm.DB) {

}

func (au Auth) DelAccount(db *gorm.DB) {

}

func (au Auth) GetAccount(db *gorm.DB) {

}
