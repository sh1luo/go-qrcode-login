package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

// 数据库连接实例
var (
	DBEngine    *gorm.DB
	RedisConn *redis.Pool
)
