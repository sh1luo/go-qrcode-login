package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/internal/service"
	"github.com/sh1luo/go-qrcode-login.git/pkg/app"
	"github.com/sh1luo/go-qrcode-login.git/pkg/errcode"
	"log"
)

func RegAuth(c *gin.Context) {
	response := app.NewResponse(c)
	params := &service.RegRequest{}
	if err := c.ShouldBind(params); err != nil {
		log.Printf("RegRequest Params Err:%s", err)
		response.ToErrResponse(errcode.InvalidParams)
	}

	svc := service.NewService(c.Request.Context())
	if err := svc.CheckAndCreate(params); err != nil {
		response.ToErrResponse(errcode.InternalErr)
	}

	response.ToResponse(gin.H{
		"code": 0,
		"msg":  "成功",
	})
}
