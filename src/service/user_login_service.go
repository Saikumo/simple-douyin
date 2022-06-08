package service

import (
	"errors"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/enum"
	"saikumo.org/simple-douyin/src/repository"
	"saikumo.org/simple-douyin/src/util"
)

type userLoginFlow struct {
	username string
	password string

	userId int64
	token  string

	response *dto.UserLoginResponse
}

func UserLogin(username string, password string) (*dto.UserLoginResponse, error) {
	return newUserLoginFlow(username, password).do()
}

func newUserLoginFlow(username string, password string) *userLoginFlow {
	return &userLoginFlow{
		username: username,
		password: password,
	}
}

func (flow *userLoginFlow) do() (*dto.UserLoginResponse, error) {
	//校验参数
	if err := flow.checkParam(); err != nil {
		return nil, err
	}
	//登录 验证用户名密码
	if err := flow.login(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()
	return flow.response, nil
}

func (flow *userLoginFlow) checkParam() error {
	//用户名为空
	if flow.username == "" {
		return errors.New(enum.USERNAME_IS_NULL)
	}
	//用户名超过最大长度
	if len(flow.username) > enum.MAX_USERNAME_LENGTH {
		return errors.New(enum.USERNAME_OVER_MAX_LENGTH)
	}
	//密码为空
	if flow.password == "" {
		return errors.New(enum.PASSWORD_IS_NULL)
	}
	//密码长度不合法
	if len(flow.password) < enum.MIN_PASSWORD_LENGTH || len(flow.password) > enum.MAX_PASSWORD_LENGTH {
		return errors.New(enum.PASSWORD_LENGTH_NOT_VALID)
	}
	return nil
}

func (flow *userLoginFlow) login() error {
	userRepository := repository.GetUserRepository()
	//验证用户名密码
	user := userRepository.FindUserByUsernameAndPassword(flow.username, flow.password)
	if user.Id == 0 {
		return errors.New(enum.USERNAME_OR_PASSWORD_IS_WRONG)
	}
	flow.userId = user.Id
	//生成token
	token, err := util.GenerateToken(user)
	if err != nil {
		return err
	}
	flow.token = token
	return nil
}

func (flow *userLoginFlow) packResponse() {
	flow.response = &dto.UserLoginResponse{
		Response: dto.Response{
			StatusCode: 0,
		},
		UserId: flow.userId,
		Token:  flow.token,
	}
}
