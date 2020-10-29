package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sh1luo/go-qrcode-login.git/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToJsonResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrResponse(err *errcode.Error) {
	if err != nil {
		r.Ctx.JSON(err.StatusCode(), err)
	}
}

//ToImgResponse does not check whether the parameter is an empty byte array.
func (r *Response) ToDataResponse(contentType string, data []byte) {
	r.Ctx.Data(http.StatusOK, contentType, data)
}

// TODO: r.ToResponseList()
func (r *Response) ToResponseList() {

}
