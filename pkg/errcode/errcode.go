package errcode

import "net/http"

type Error struct {
	Code uint     `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

func NewErr(code uint, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) GetCode() uint {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case InvalidParams.GetCode():
		return http.StatusBadRequest
	case InternalErr.GetCode():
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
