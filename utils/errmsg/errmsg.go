package errmsg

const (
	SUCCESS = 200
	ERROR = 500

	// code = 1000... 用户模块错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST = 1004
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_WRONG = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT = 1008

	// code = 2000... 文章模块错误
	ERROR_ART_NOT_EXIST = 2001

	// code = 3000... 分类模块错误
	ERROR_CARTNAME_USERD = 3001
	ERROR_CAR_NOT_EXIST = 3002

)

var codeMsg = map[int] string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已经存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已经过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT: "用户无权限",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CARTNAME_USERD: "分类名已存在",
	ERROR_CAR_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
