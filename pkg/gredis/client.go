package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"log"
)

//Set use custom data structure to
func Set(key string, data interface{}, expiration int) error {
	conn := global.RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value, "EX", expiration)
	if err != nil {
		return err
	}

	return nil
}

//Get
func Get(key string) *chain {
	conn := global.RedisConn.Get()
	defer conn.Close()

	d, err := redis.Bytes(conn.Do("GET", key))
	return &chain{
		data: d,
		err:  err,
	}
}

func GetAnyInt(key string) int {
	conn := global.RedisConn.Get()
	defer conn.Close()

	d, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		log.Printf("redis.Int err: %s", err)
		return -1
	}
	return d
}

func Exist(key string) bool {
	conn := global.RedisConn.Get()
	defer conn.Close()

	b, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return b
}

func TTL(key string) int {
	conn := global.RedisConn.Get()
	defer conn.Close()

	t, err := redis.Int(conn.Do("TTL", key))
	if err != nil {
		return 0
	}
	return t
}
