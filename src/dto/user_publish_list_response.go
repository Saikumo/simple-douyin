package dto

type UserPublishListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}
