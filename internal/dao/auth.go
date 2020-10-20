package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sh1luo/go-qrcode-login.git/internal/model"
	"time"
)

func (d *Dao) CreateAccount(appKey, appSecret string) error {
	account := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
		Model:     gorm.Model{},
	}
	return account.Create(d.engine)
}

func (d *Dao) UpdateAccount(id uint, lastLoginDate time.Time, LastLoginIp string) error {
	account := model.Auth{
		Model: gorm.Model{
			ID: id,
		},
	}
	updMap := make(map[string]interface{}, 2)
	updMap["last_login_date"] = lastLoginDate
	updMap["login_ip"] = LastLoginIp
	return account.UpdateAccount(d.engine, updMap)
}

func (d *Dao) DeleteAccount(id uint) error {
	account := model.Auth{
		Model: gorm.Model{
			ID: id,
		},
	}
	return account.DelAccount(d.engine)
}
