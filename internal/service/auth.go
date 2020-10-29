package service

import (
	"encoding/base64"
	"github.com/sh1luo/go-qrcode-login.git/pkg/errcode"
	"golang.org/x/crypto/scrypt"
)

type RegRequest struct {
	AppKey    string `json:"app_key" form:"app_key" validate:"gte=3,lte=15"`
	AppSecret string `json:"app_secret" form:"app_secret" validate:"gte=3,lte=15"`
}

//CheckExit returns 3 results,already/not registered and ServerInternalErr.
func (svc *Service) CheckAuthExit(regInfo *RegRequest) *errcode.Error {
	au, err := svc.Dao.GetAccountByKey(regInfo.AppKey)
	if err != nil {
		return errcode.ServerInternalErr
	}

	if au.ID > 0 {
		return errcode.AppNameHasExit
	}

	return errcode.AppNameHasNotExit
}

// Create is the main business logic of the registration part,
// which returns nil when all processes exec successfully,
// custom and specific Error type when has any error and
// you can handle according to different error returns.
func (svc *Service) CreateAuth(regInfo *RegRequest) *errcode.Error {
	returnErr := svc.CheckAuthExit(regInfo)
	if returnErr == errcode.AppNameHasNotExit {
		encodePasswd, err := handlePlainPasswd([]byte(regInfo.AppSecret))
		if err != nil {
			return errcode.ServerInternalErr
		}

		err = svc.Dao.CreateAccount(regInfo.AppKey, encodePasswd)
		if err != nil {
			return errcode.ServerInternalErr
		}
		return nil
	}

	return returnErr
}

//GetAuth is the main business logic of the login part
//by AppKey and AppSecret.
func (svc *Service) GetAuth(regInfo *RegRequest) (uint, *errcode.Error) {
	au, err := svc.Dao.GetAccountByKey(regInfo.AppKey)
	if err != nil {
		return 0, errcode.ServerInternalErr
	}

	var curEncodePasswd string
	curEncodePasswd, err = handlePlainPasswd([]byte(regInfo.AppSecret))
	if au.ID <= 0 || au.AppSecret != curEncodePasswd {
		return 0, errcode.AppAuthFailed
	}

	return au.ID, nil
}

func handlePlainPasswd(plainPasswd []byte) (string, error) {
	// salt is a random string,such as salt = "O@S#S$A%"
	var salt = []byte{0x4f, 0x40, 0x53, 0x23, 0x53, 0x24, 0x41, 0x25}
	scryptPasswd, err := scrypt.Key(plainPasswd, salt, 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	afterBase64 := base64.StdEncoding.EncodeToString(scryptPasswd)
	return afterBase64, nil
}
