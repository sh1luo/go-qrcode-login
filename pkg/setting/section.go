package setting

import (
	"log"
	"time"
)

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MySQLSettings struct {
	Host         string
	Port         string
	DBName       string
	Username     string
	Password     string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettings struct {
	Host     string
	Port     string
	DB       int8
	Password string
}

type JwtSettings struct {
	Secret string
	Iss string
	Expire time.Duration
}

func (st *Setting) ReadSection(key string, value interface{}) error {
	if err := st.vp.UnmarshalKey(key, value); err != nil {
		log.Printf("%s:%s read config struct err:%s", key, value, err)
		return err
	}
	return nil
}
