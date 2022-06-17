package service

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/repository"
	"strconv"
	"time"
)

type videoFeedFlow struct {
	latestTimeStr string //纳秒 字符串

	latestTime time.Time
	nextTime   time.Time
	videoArray []*dto.Video

	response *dto.VideoFeedResponse
}

func VideoFeed(latestTimeStr string) (*dto.VideoFeedResponse, error) {
	common.Logger.Infof("视频流推送")
	return newVideoFeedFlow(latestTimeStr).do()
}

func newVideoFeedFlow(latestTimeStr string) *videoFeedFlow {
	return &videoFeedFlow{
		latestTimeStr: latestTimeStr,
	}
}

func (flow *videoFeedFlow) do() (*dto.VideoFeedResponse, error) {
	//校验参数
	if err := flow.checkParam(); err != nil {
		return nil, err
	}
	//获取推送内容
	if err := flow.feed(); err != nil {
		return nil, err
	}
	//打包返回值
	flow.packResponse()

	return flow.response, nil
}

func (flow *videoFeedFlow) checkParam() error {
	//如果时间戳为空则为当前时间
	if flow.latestTimeStr == "" || flow.latestTimeStr == "0" {
		flow.latestTime = time.Now()
	} else {
		latestTimeUnixMilli, err := strconv.ParseInt(flow.latestTimeStr, 10, 64)
		if err != nil {
			return err
		}
		flow.latestTime = time.Unix(0, latestTimeUnixMilli*1e6)
	}
	return nil
}

func (flow *videoFeedFlow) feed() error {
	videoRepository := repository.GetVideoRepository()
	userRepository := repository.GetUserRepository()
	//拿出视频
	videoArray, err := videoRepository.FindVideoByLimitAndLatestTime(10, flow.latestTime)
	if err != nil {
		return err
	}

	for _, video := range videoArray {
		user := userRepository.FindUserByUserId(video.UserId)
		//用户不存在
		if user.ID == 0 {
			return common.UserNotExistError
		}

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
		flow.nextTime = video.CreatedAt
	}

	return nil
}

func (flow *videoFeedFlow) packResponse() {
	flow.response = &dto.VideoFeedResponse{
		Response: dto.Response{
			StatusCode: common.SUCCESS_STATUS_CODE,
			StatusMsg:  "",
		},
		NextTime:  flow.nextTime.Unix(),
		VideoList: flow.videoArray,
	}
}
