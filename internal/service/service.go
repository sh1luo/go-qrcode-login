package service

import (
	"context"
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