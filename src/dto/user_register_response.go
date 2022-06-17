package dto

type UserRegisterResponse struct {
	Response
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}
