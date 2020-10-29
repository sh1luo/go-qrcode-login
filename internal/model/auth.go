package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Auth struct {
	AppKey        string     `json:"app_key"`
	AppSecret     string     `json:"app_secret"`
	LastLoginDate *time.Time `json:"last_login_date"`
	LastLoginIp   string     `json:"login_ip"`
	gorm.Model
}

func (au Auth) TableName() string {
	return "auth"
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

func (au *Auth) GetAccountByKey(db *gorm.DB) (at Auth, err error) {
	err = db.Find(&at, "app_key = ?", au.AppKey).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return at, err
	}
	return at, nil
}

func (au *Auth) GetAccountByID(db *gorm.DB) (at Auth, err error) {
	err = db.Find(&at, "id = ?", au.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return at, err
	}
	return at, nil
}