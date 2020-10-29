package service

import "time"

//TODO Provide CACHE interface
//Cache
type Cache interface {
	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) interface{}
	Delete(key string) error
	CacheID() string
}
