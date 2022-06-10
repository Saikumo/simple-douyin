package service

import "saikumo.org/simple-douyin/src/dto"

type videoFeedFlow struct {
	latestTime int64
}

func VideoFeed(latestTime int64) (*dto.VideoFeedResponse, error) {
	return newVideoFeedFlow(latestTime).do()
}

func newVideoFeedFlow(latestTime int64) *videoFeedFlow {
	return &videoFeedFlow{
		latestTime: latestTime,
	}
}

func (flow *videoFeedFlow) do() (*dto.VideoFeedResponse, error) {
	return nil, nil
}
