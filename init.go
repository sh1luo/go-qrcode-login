package main

import (
	"fmt"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/internal/model"
	setting2 "github.com/sh1luo/go-qrcode-login.git/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		fmt.Printf("Read config file section err:%s", err)
	}
	if err := setupDB(); err != nil {
		fmt.Printf("Get DBEngine err:%s", err)
	}
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		panic(fmt.Errorf("Read config file err: %s\n", err))
	}

	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		panic(fmt.Errorf("Read config file:Server section err: %s\n", err))
	}
	if err = setting.ReadSection("MySQL", &global.MySQLSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Redis", &global.RedisSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("JWT", &global.JwtSetting); err != nil {
		return err
	}

	return nil
}

func setupDB() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.MySQLSetting)
	if err != nil {
		return err
	}
	return nil
}
