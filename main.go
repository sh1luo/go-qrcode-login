package main

import (
	"errors"
	"fmt"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/internal/model"
	"github.com/sh1luo/go-qrcode-login.git/internal/router"
	"github.com/sh1luo/go-qrcode-login.git/pkg/setting"
	"log"
	"net/http"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Printf("Read config file section err:%s", err)
	}

	err = setupDB()
	if err != nil {
		log.Printf("Get DBEngine err:%s", err)
	}
}

func setupSetting() error {
	globalSetting, err := setting.NewSetting()
	if err != nil {
		panic(fmt.Errorf("Read config file err: %s\n", err))
	}

	err = globalSetting.ReadSection("Server", global.ServerSetting)
	if err != nil {
		panic(fmt.Errorf("Read config file:Server section err: %s\n", err))
	}

	err = globalSetting.ReadSection("MySQL", global.MySQLSetting)
	if err != nil {
		return err
	}

	err = globalSetting.ReadSection("Redis", global.RedisSetting)
	if err != nil {
		return err
	}

	err = globalSetting.ReadSection("JWT", global.JwtSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDB() error {
	fmt.Println("setupDb starting...")

	var err error
	global.DBEngine, err = model.NewDBEngine(global.MySQLSetting)
	if err != nil {
		return err
	}
	if global.RdbEngine = model.NewRedisEngine(global.RedisSetting); global.RdbEngine != nil {
		return errors.New("redis connected err")
	}
	return nil
}

func main() {
	//gin.SetMode(global.ServerSetting.RunMode)
	route := router.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        route,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("mysql password: ", global.MySQLSetting.Password)

	if err := s.ListenAndServe(); err != nil {
		log.Printf("app start err: %s", err)
	}
}
