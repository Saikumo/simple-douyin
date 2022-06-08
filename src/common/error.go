package common

import "errors"

var (
	UsernameIsNullError            = errors.New("用户名不能为空")
	UsernameOverMaxLengthError     = errors.New("用户名长度不能超过32个字符")
	PasswordIsNullError            = errors.New("密码不能为空")
	PasswordLengthNotValidError    = errors.New("密码长度必须为6到32个字符")
	UsernameAlreadyExistError      = errors.New("用户名已经存在")
	UsernameOrPasswordIsWrongError = errors.New("用户名或密码错误")
	TokenIsNotValid                = errors.New("token不正确")
	TokenIsExpired                 = errors.New("token已过期")
	TokenIsNotExist                = errors.New("token不存在")
	UserIsNotExist                 = errors.New("用户不存在")
)
