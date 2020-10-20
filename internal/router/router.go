package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/internal/router/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/reg", api.RegAuth)
	//apis1 := r.Group("api/v1")
	//{
	//
	//}

	return r
}
