package util

import (
	"fmt"
	"saikumo.org/simple-douyin/src/entity"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := &entity.User{
		Username: "233333",
		Password: "233333",
	}
	token, err := GenerateToken(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
