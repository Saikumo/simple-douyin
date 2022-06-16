package common

import (
	"errors"
	"fmt"
)

var (
	UsernameIsNullError            = errors.New("用户名不能为空")
	UsernameOverMaxLengthError     = errors.New(fmt.Sprintf("用户名长度不能超过%d个字符", MAX_USERNAME_LENGTH))
	PasswordIsNullError            = errors.New("密码不能为空")
	PasswordLengthNotValidError    = errors.New(fmt.Sprintf("密码长度必须为%d到%d个字符", MIN_PASSWORD_LENGTH, MAX_PASSWORD_LENGTH))
	UsernameAlreadyExistError      = errors.New("用户名已经存在")
	UsernameOrPasswordIsWrongError = errors.New("用户名或密码错误")
	TokenIsNotValidError           = errors.New("token不正确")
	TokenIsExpiredError            = errors.New("token已过期")
	TokenIsNotExistError           = errors.New("token不存在")
	UserIsNotExistError            = errors.New("用户不存在")
	VideoTitleIsNullError          = errors.New("视频标题不能为空")
	VideoTitleOverMaxLengthError   = errors.New(fmt.Sprintf("视频标题不能超过%d个字符", MAX_VIDEO_TITLE_LENGTH))
	UnsupportedVideoTypeError      = errors.New("不支持的视频文件类型")
)
