package dto

type UserInfo struct {
	Id       int64
	Username string `json:"name"`
}

type UserInfoResponse struct {
	*Response
	UserInfo *UserInfo `json:"user"`
}
