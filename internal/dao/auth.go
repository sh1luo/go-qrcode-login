package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sh1luo/go-qrcode-login.git/internal/model"
)

func (d *Dao) CreateAccount(appKey,appSecret,salt string) error {
	account := model.Auth{
		AppKey:        appKey,
		AppSecret:     appSecret,
		Salt:          salt,
		Model:         gorm.Model{},
	}
	return account.Create(d.engine)
}