package router

import "github.com/gin-gonic/gin"

func NewRouter(ctx *gin.Context) *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("/")
	}

	return r
}