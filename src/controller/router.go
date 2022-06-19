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

	feedRouter := douyinRouter.Group("/feed")
	feedRouter.GET("/", VideoFeed)

	publishRouter := douyinRouter.Group("/publish")
	publishRouter.POST("/action/", middleware.JwtAuth(), UserPublish)
	publishRouter.GET("/list/", UserPublishList)

	favoriteRouter := douyinRouter.Group("/favorite")
	favoriteRouter.POST("/action/", middleware.JwtAuth(), UserFavorite)

	addr := strconv.Itoa(common.Config.GetInt("server.port"))
	if err := r.Run(":" + addr); err != nil {
		panic(err)
	}
}
