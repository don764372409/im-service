package utils

import (
	"fmt"
)

var (
	ERR_USER_REGISTER_USERNAME_NULL = NewError(1000, "username is null", "账号不能为空")
	ERR_USER_REGISTER_PASSWORD_NULL = NewError(1001, "password is null", "密码不能为空")
	//账号长度为6-12位
	ERR_USER_REGISTER_USERNAME_LENGTH = NewError(1002, "username length is error", "账号长度为6-20位")
	ERR_USER_REGISTER_PASSWORD_LENGTH = NewError(1003, "password length is error", "密码长度为6-20位")
	ERR_USER_NOT_EXIST                = NewError(1004, "user not exist", "用户不存在")

	ERR_UPLOAD_FILE        = NewError(1005, "upload file error", "上传文件失败")
	ERR_PARAM_PARSE        = NewError(1006, "param parse error", "参数解析失败")
	ERR_GET_USER_INFO      = NewError(1007, "get user info error", "获取用户信息失败")
	ERR_LOGIN_FAIL         = NewError(1008, "login fail", "登录失败,请刷新失败")
	ERR_REGISTER_FAIL      = NewError(1009, "register fail", "注册失败")
	ERR_SECRET_UPDATE_FAIL = NewError(1010, "secret update fail", "秘钥修改失败")
)

type Error struct {
	Code       int
	Msg        string
	MsgZh      string
	IsHasStack bool
}

func NewError(errCode int, msg, msgZh string) *Error {
	return &Error{
		Code:  errCode,
		Msg:   msg,
		MsgZh: msgZh,
	}
}

func NewSysError(err error) *Error {
	return &Error{
		Code:  500,
		Msg:   err.Error(),
		MsgZh: err.Error(),
	}
}

func (e *Error) ErrorCode() int {
	return e.Code
}

func (e *Error) Error() string {
	return fmt.Sprintf("-> Code:%d, Msg:%s, MsgZh:%s", e.Code, e.Msg, e.MsgZh)
}