package service

import (
	"errors"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/entity"
	"saikumo.org/simple-douyin/src/enum"
	"saikumo.org/simple-douyin/src/repository"
	"saikumo.org/simple-douyin/src/util"
)

type userRegisterFlow struct {
	username string
	password string

	userId int64
	token  string

	response *dto.UserRegisterResponse
}

func UserRegister(username string, password string) (*dto.UserRegisterResponse, error) {
	return newUserRegisterFlow(username, password).do()
}

func newUserRegisterFlow(username string, password string) *userRegisterFlow {
	return &userRegisterFlow{
		username: username,
		password: password,
	}
}

func (flow *userRegisterFlow) do() (*dto.UserRegisterResponse, error) {
	//校验参数
	if err := flow.checkParam(); err != nil {
		return nil, err
	}
	//注册
	if err := flow.register(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()
	return flow.response, nil
}

func (flow *userRegisterFlow) checkParam() error {
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

func (flow *userRegisterFlow) register() error {
	userRepository := repository.GetUserRepository()
	//检测用户是否存在
	if userRepository.IsUserExistByUsername(flow.username) {
		return errors.New(enum.USERNAME_ALREADY_EXIST)
	}
	//创建用户
	user := entity.User{Username: flow.username, Password: flow.password}
	if err := userRepository.CreateUser(&user); err != nil {
		return err
	}
	flow.userId = user.Id
	//生成token
	token, err := util.GenerateToken(&user)
	if err != nil {
		return err
	}
	flow.token = token
	return nil
}

func (flow *userRegisterFlow) packResponse() {
	flow.response = &dto.UserRegisterResponse{
		Response: dto.Response{
			StatusCode: 0,
		},
		UserId: flow.userId,
		Token:  flow.token,
	}
}
