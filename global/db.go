package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

// 数据库连接实例
var (
	DBEngine    *gorm.DB
	RedisEngine *redis.Client
)
