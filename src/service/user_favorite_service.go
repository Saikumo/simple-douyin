package service

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/repository"
)

type userFavoriteFlow struct {
	userId     uint
	videoId    uint
	actionType uint

	response *dto.Response
}

func UserFavorite(userId, videoId, actionType uint) (*dto.Response, error) {
	return newUserFavoriteFlow(userId, videoId, actionType).do()
}

func newUserFavoriteFlow(userId, videoId, actionType uint) *userFavoriteFlow {
	return &userFavoriteFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: actionType,
	}
}

func (flow *userFavoriteFlow) do() (*dto.Response, error) {
	//校验参数
	if err := flow.checkParam(); err != nil {
		return nil, err
	}
	//点赞或取消点赞
	if err := flow.userFavorite(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()

	return flow.response, nil
}

func (flow *userFavoriteFlow) checkParam() error {
	//actionType必须是1或者2 1点赞 2取消点赞
	if flow.actionType != 1 && flow.actionType != 2 {
		return common.UnsupportedUserFavoriteActionTypeError
	}
	return nil
}

func (flow *userFavoriteFlow) userFavorite() error {
	favoriteRepository := repository.GetFavoriteRepository()
	var err error

	switch flow.actionType {
	//点赞
	case 1:
		err = favoriteRepository.UserFavorite(flow.userId, flow.videoId)
	//取消点赞
	case 2:
		err = favoriteRepository.UserUndoFavorite(flow.userId, flow.videoId)
	}
	if err != nil {
		return err
	}

	return nil
}

func (flow *userFavoriteFlow) packResponse() {
	flow.response = &dto.Response{
		StatusCode: common.SUCCESS_STATUS_CODE,
	}
}
