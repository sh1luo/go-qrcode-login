package errcode

import "net/http"

type Error struct {
	Code uint     `json:"code"`
	Msg  string   `json:"msg,omitempty"`
	Data []string `json:"details,omitempty"`
}

//NewErr return a new custom error type,which has zero or some details info.
func NewErr(code uint, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// func (e Error) AddParams(key string,value interface{}) Error {
func (e Error) AddParams(params ...string) Error {
	e.Data = params
	return e
}

func (e Error) Error() string {
	return e.Msg
}

func (e *Error) GetCode() uint {
	return e.Code
}

func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case InvalidParams.GetCode():
		return http.StatusOK
	case ServerInternalErr.GetCode():
		return http.StatusInternalServerError
	default:
		return http.StatusOK
	}
}
