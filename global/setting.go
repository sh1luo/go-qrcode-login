package global

import "github.com/sh1luo/go-qrcode-login.git/pkg/setting"

var (
	ServerSetting *setting.ServerSettings
	MySQLSetting  *setting.MySQLSettings
	RedisSetting  *setting.RedisSettings
	JwtSetting    *setting.JwtSettings
)
