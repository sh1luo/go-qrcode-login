package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func EncodePasswd(passwd []byte) {
	p, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword(p, []byte("shiluo")); err != nil {
		fmt.Println(err)
	}
	fmt.Println("correct")
}
