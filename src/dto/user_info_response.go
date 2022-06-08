package dto

import "saikumo.org/simple-douyin/src/entity"

type UserInfoResponse struct {
	*Response
	User *entity.User
}
