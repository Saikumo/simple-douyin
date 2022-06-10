package controller

import (
	"github.com/gin-gonic/gin"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/middleware"
	"strconv"
)

func InitRouter() {
	r := gin.Default()

	r.Static("/static/", "./static")

	douyinRouter := r.Group("/douyin")

	userRouter := douyinRouter.Group("/user")
	userRouter.POST("/register/", UserRegister)
	userRouter.POST("/login/", UserLogin)
	userRouter.GET("/", middleware.JwtAuth(), UserInfo)

	addr := strconv.Itoa(common.Config.GetInt("server.port"))
	if err := r.Run(":" + addr); err != nil {
		panic(err)
	}
}
