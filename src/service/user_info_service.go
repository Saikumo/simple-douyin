package service

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/entity"
	"saikumo.org/simple-douyin/src/repository"
)

type userInfoFlow struct {
	userId uint

	user *entity.User

	response *dto.UserInfoResponse
}

func UserInfo(userId uint) (*dto.UserInfoResponse, error) {
	common.Logger.Infof("获取用户信息，用户id为%d", userId)
	return newUserInfoFlow(userId).do()
}

func newUserInfoFlow(userId uint) *userInfoFlow {
	return &userInfoFlow{
		userId: userId,
	}
}

func (flow *userInfoFlow) do() (*dto.UserInfoResponse, error) {
	//查询用户
	if err := flow.info(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()
	return flow.response, nil
}

func (flow *userInfoFlow) info() error {
	userRepository := repository.GetUserRepository()
	user := userRepository.FindUserByUserId(flow.userId)
	if user.ID == 0 {
		return common.UserIsNotExistError
	}
	flow.user = user
	return nil
}

func (flow *userInfoFlow) packResponse() {
	flow.response = &dto.UserInfoResponse{
		Response: &dto.Response{
			StatusCode: common.SUCCESS_STATUS_CODE,
		},
		UserInfo: &dto.UserInfo{
			Id:       flow.user.ID,
			Username: flow.user.Username,
		},
	}
}
