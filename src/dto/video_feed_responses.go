package dto

type VideoFeedResponse struct {
	*Response
	NextTime  uint
	VideoList []*Video
}
