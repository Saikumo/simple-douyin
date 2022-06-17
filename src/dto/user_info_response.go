package dto

type UserInfoResponse struct {
	Response
	UserInfo *UserInfo `json:"user"`
}
