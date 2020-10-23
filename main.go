package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/global"
	_ "github.com/sh1luo/go-qrcode-login.git/global"
	"github.com/sh1luo/go-qrcode-login.git/internal/router"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	route := router.NewRouter()
	s := http.Server{
		Addr:    ":" + global.ServerSetting.HttpPort,
		Handler: route,
		//ReadTimeout:    global.ServerSetting.ReadTimeout,
		//WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("app start err: %s", err)
	}
}
