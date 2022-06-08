package service

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/repository"
	"testing"
)

func TestUserRegister(t *testing.T) {
	common.InitConfigByPath("../../")
	repository.InitDB()

	if _, err := UserRegister("222", "222222"); err != nil {
		panic(err)
	}
}
