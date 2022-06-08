package service

import (
	"saikumo.org/simple-douyin/src/config"
	"saikumo.org/simple-douyin/src/repository"
	"testing"
)

func TestUserRegister(t *testing.T) {
	config.InitConfigByPath("../../")
	repository.InitDB()

	if _, err := UserRegister("222", "222222"); err != nil {
		panic(err)
	}
}
