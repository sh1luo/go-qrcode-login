package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {vp *viper.Viper}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs")
	vp.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal err: viper read config file err:%s",err))
	}
	return &Setting{vp: vp}, nil
}
