package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/internal/dao"
)

type Service struct {
	Ctx context.Context
	Dao *dao.Dao
}

func NewService(context context.Context) Service {
	svc := Service{Ctx: context}
	svc.Dao = dao.NewDao(global.DBEngine)
	return svc
}

// TODO import!!!!****Change the parameter list to the interface type*****
//AddToCache is a general service method, used to add specified content to the CACHE(redis,memcached,etc), which can be used by all services.
//func (svc *Service) AddToCache(val interface{})  errcode.Error {
//	// generate token by appKey.
//	//var uuidToken []byte
//	//src := util.Base64EncodeToString([]byte(qcStr))
//	//uuidToken = append(uuidToken, []byte(src)[:8]...)
//	//uuidToken = append(uuidToken, []byte{'=', '='}...)
//	uuidToken := svc.TokenID(10)
//
//	// ADD token:userid TO CACHE
//	err := gredis.Set(svc.Ctx, uuidToken, val, global.JwtSetting.Expire)
//	if err != nil {
//		e := errcode.ServerInternalErr.AddParams("set key-value cache err")
//		return  e
//	}
//
//	return  errcode.Success.AddParams(uuidToken)
//}

func (svc *Service) TokenID(l uint8) string {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
