package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_USER:               "已存在该用户名称",
	ERROR_NOT_EXIST_USER:           "该用户不存在",
	ERROR_EXIST_TASKGROUP:          "已存在该分组名称",
	ERROR_NOT_EXIST_TASKGROUP:      "该分组不存在",
	ERROR_EXIST_TASK:               "已存在该任务名称",
	ERROR_NOT_EXIST_TASK:           "该任务不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_LOGIN:                    "登录失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
