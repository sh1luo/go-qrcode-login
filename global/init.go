package global

import (
	"errors"
	"fmt"
	"github.com/sh1luo/go-qrcode-login.git/internal/model"
	"github.com/sh1luo/go-qrcode-login.git/pkg/setting"
	"github.com/sh1luo/go-qrcode-login.git/pkg/validator"

	"log"
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

	Validator = validator.NewCustomValidator()
}

func setupSetting() error {
	globalSetting, err := setting.NewSetting()
	if err != nil {
		panic(fmt.Errorf("Read config file err: %s\n", err))
	}

	err = globalSetting.ReadSection("Server", &ServerSetting)
	if err != nil {
		panic(fmt.Errorf("Read config file:Server section err: %s\n", err))
	}

	err = globalSetting.ReadSection("MySQL", &MySQLSetting)
	if err != nil {
		return err
	}

	err = globalSetting.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}

	err = globalSetting.ReadSection("JWT", &JwtSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDB() error {
	//fmt.Println("setupDb starting...")

	var err error
	DBEngine, err = model.NewDBEngine(MySQLSetting)
	if err != nil {
		return err
	}
	if RedisEngine = model.NewRedisEngine(RedisSetting); RedisEngine != nil {
		return errors.New("redis connected err")
	}
	return nil
}
