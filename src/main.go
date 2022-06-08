package main

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/controller"
	"saikumo.org/simple-douyin/src/repository"
)

func main() {
	//初始化
	common.InitConfig()
	common.InitLogger()
	repository.InitDB()
	controller.InitRouter()
}
