package util

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// 获取随机字符串
func GetRandomString(l int) string {
	str := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, str[r.Intn(len(str))])
	}
	return string(result)
}

func Base64EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
