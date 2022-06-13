package dto

type UserLoginResponse struct {
	*Response
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}
