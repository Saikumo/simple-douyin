package dto

type VideoFeedResponse struct {
	Response
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list"`
}
