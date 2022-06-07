package main

import (
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/config"
	"saikumo.org/simple-douyin/src/repository"
	"saikumo.org/simple-douyin/src/router"
)

func main() {
	//初始化
	config.InitConfig()
	common.InitLogger()
	repository.InitDB()
	router.InitRouter()
}
