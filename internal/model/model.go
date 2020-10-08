package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sh1luo/go-qrcode-login.git/pkg/setting"
)

func NewDBEngine(dbSetting setting.MySQLSettings) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
			dbSetting.Username,
			dbSetting.Password,
			dbSetting.Host,
			dbSetting.Port,
			dbSetting.DBName,
			dbSetting.Charset),
		dbSetting.ParseTime)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db, nil
}
