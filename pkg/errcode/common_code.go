package errcode

var (
	Success     *Error = NewErr(0, "成功")
	InternalErr        = NewErr(1, "服务器内部错误")

	// Register
	InvalidParams = NewErr(100005, "注册所需参数错误")
)
