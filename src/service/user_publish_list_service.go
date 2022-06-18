package service

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/repository"
)

type userPublishListFlow struct {
	userId uint

	videoArray []*dto.Video

	response *dto.UserPublishListResponse
}

func UserPublishList(userId uint) (*dto.UserPublishListResponse, error) {
	return newUserPublishListFlow(userId).do()
}

func newUserPublishListFlow(userId uint) *userPublishListFlow {
	return &userPublishListFlow{
		userId: userId,
	}
}

func (flow *userPublishListFlow) do() (*dto.UserPublishListResponse, error) {
	//获取用户发布列表
	if err := flow.userPublishList(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()

	return flow.response, nil
}

func (flow *userPublishListFlow) userPublishList() error {
	userRepository := repository.GetUserRepository()
	user := userRepository.FindUserByUserId(flow.userId)
	//用户不存在
	if user.ID == 0 {
		return common.UserNotExistError
	}

	videoRepository := repository.GetVideoRepository()
	videoArray, err := videoRepository.FindVideoByUserId(flow.userId)
	if err != nil {
		return err
	}

	for _, video := range videoArray {
		flow.videoArray = append(flow.videoArray, &dto.Video{
			Id: video.ID,
			UserInfo: &dto.UserInfo{
				Id:       user.ID,
				Username: user.Username,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			Title:         video.Title,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		})
	}
	return nil
}

func (flow *userPublishListFlow) packResponse() {
	flow.response = &dto.UserPublishListResponse{
		Response: dto.Response{
			StatusCode: common.SUCCESS_STATUS_CODE,
		},
		VideoList: flow.videoArray,
	}
}
