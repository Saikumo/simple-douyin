package repository

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/entity"
	"testing"
)

func TestDB(t *testing.T) {
	common.InitConfigByPath("../../")
	InitDB()

	user := entity.User{
		Username: "233",
		Password: "233",
	}

	DB.Save(&user)
}
