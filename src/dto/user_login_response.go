package dto

type UserLoginResponse struct {
	*Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
