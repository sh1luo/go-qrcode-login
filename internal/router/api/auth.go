package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/internal/service"
	"github.com/sh1luo/go-qrcode-login.git/pkg/app"
	"github.com/sh1luo/go-qrcode-login.git/pkg/errcode"
	"github.com/sh1luo/go-qrcode-login.git/pkg/gredis"
	"log"
)

// RegAuth realize parameter binding, field verification,
// and execute judgment logic, choose whether to create a user
// according to the returned result.
func RegAuth(c *gin.Context) {
	var params service.RegRequest
	var err error
	response := app.NewResponse(c)

	if err = c.ShouldBind(&params); err != nil {
		response.ToErrResponse(errcode.ServerInternalErr)
		return
	}

	err = global.Validator.ValidateStruct(params)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams)
		return
	}

	svc := service.NewService(c.Request.Context())
	if resp := svc.CreateAuth(&params); resp != nil {
		response.ToErrResponse(resp)
		return
	}

	response.ToJsonResponse(errcode.AppRegSuccess)
}

//GetAuth login by AppKey and AppSecret.
func GetAuth(c *gin.Context) {
	var err error
	response := app.NewResponse(c)
	params := service.RegRequest{}

	if err = c.ShouldBind(&params); err != nil {
		response.ToErrResponse(errcode.ServerInternalErr)
		return
	}

	err = global.Validator.ValidateStruct(&params)
	if err != nil {
		response.ToErrResponse(errcode.InvalidParams)
		return
	}

	svc := service.NewService(c.Request.Context())
	uid, e := svc.GetAuth(&params)
	if e != nil {
		response.ToErrResponse(e)
		return
	}

	// TODO NEED A CACHE INTERFACE
	token := svc.TokenID(10)
	err = gredis.Set(token, uid, global.JwtSetting.AppExpire)
	if err != nil {
		response.ToErrResponse(errcode.ServerInternalErr)
		log.Printf("redis写入失败:%s\n", err)
		return
	}

	response.ToJsonResponse(errcode.Success.AddParams(token))
}

func TestFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "成功！",
	})
}
