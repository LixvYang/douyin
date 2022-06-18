// Package errmsg provides error message package.
package errmsg

var (
	OK                  = &Errmsg{StatusCode: 0, StatusMsg: "OK"}
	ERROR               = &Errmsg{StatusCode: 1, StatusMsg: "ERROR"}
	InternalServerError = &Errmsg{StatusCode: 10001, StatusMsg: "网络错误"}

	ErrBind     = &Errmsg{StatusCode: 20001, StatusMsg: "参数错误"}
	ErrDatabase = &Errmsg{StatusCode: 20002, StatusMsg: "数据库错误"}
	ErrToken    = &Errmsg{StatusCode: 20003, StatusMsg: "Token错误"}
	ErrGetFile  = &Errmsg{StatusCode: 20004, StatusMsg: "获取文件错误"}
	ErrFileType = &Errmsg{StatusCode: 20005, StatusMsg: "文件类型错误"}
	ErrSaveFile = &Errmsg{StatusCode: 20006, StatusMsg: "保存文件错误"}

	ErrUserNotFound  = &Errmsg{StatusCode: 20002, StatusMsg: "用户不存在"}
	ErrUsernameExist = &Errmsg{StatusCode: 20003, StatusMsg: "用户名已存在"}
	ErrUserRegister  = &Errmsg{StatusCode: 20004, StatusMsg: "用户注册失败"}
	ErrUserLogin     = &Errmsg{StatusCode: 20005, StatusMsg: "用户未登录"}
)
