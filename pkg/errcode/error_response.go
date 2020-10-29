package errcode

var (
	// Common status
	Success           = NewErr(0, "操作成功")
	ServerInternalErr = NewErr(500000, "服务器内部错误")
	InvalidParams     = NewErr(400000, "请求所需参数错误")

	// Auth status
	AppRegSuccess     = NewErr(100000, "注册成功")
	AppLoginSuccess   = NewErr(100001, "登陆成功")
	AppNameHasExit    = NewErr(100002, "用户名已存在")
	AppNameHasNotExit = NewErr(100003, "用户不存在")
	AppAuthFailed     = NewErr(100004, "用户名密码验证失败")
	AppFailed         = NewErr(100005, "验证失败")

	// QrCode status
	QrCodeExist   = NewErr(110000, "") // 未过期
	QrCodeExpired = NewErr(110001, "") // 已过期
	QrCodeScanned = NewErr(110002, "") // 已被扫描
)

var (

)

// what i want?
// such as this ?
/*
	response.ToSuccessWithParams(params map[string]interface{}) {
		r.Ctx.Json(http.StatusOK,)
	}

*/
