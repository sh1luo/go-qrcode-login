package setting

import (
	"log"
	"time"
)

type ServerSettings struct {
	RunMode      string
	Addr         string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MySQLSettings struct {
	Host         string
	DBName       string
	Username     string
	Password     string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettings struct {
	Host      string
	DB        int8
	Password  string
	MaxIdle   int
	MaxActive int
}

type JwtSettings struct {
	Secret       string
	Iss          string
	AppExpire    int
	QrCodeExpire int
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		log.Printf("%s:%s read config struct err:%s", k, v, err)
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
