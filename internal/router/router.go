package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/internal/router/api"
	v1 "github.com/sh1luo/go-qrcode-login.git/internal/router/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/reg", api.RegAuth)
	r.POST("/login", api.GetAuth)

	r.Group("/qrcode")
	{
		r.GET("/qrcode", v1.GetQrcode)
		r.GET("/status", v1.GetQrCodeStatus)
		r.POST("/l/:qrcode", v1.LoginByQrcode)
	}

	r.GET("/ping", api.TestFunc)
	return r
}
