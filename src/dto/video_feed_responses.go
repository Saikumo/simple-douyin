package dto

type VideoFeedResponse struct {
	*Response
	NextTime  int64
	VideoList []*Video
}
