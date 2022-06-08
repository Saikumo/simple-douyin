package common

import (
	"github.com/gin-gonic/gin"
	"saikumo.org/simple-douyin/src/controller"
	"strconv"
)

func InitRouter() {
	r := gin.Default()

	douyinRouter := r.Group("/douyin")

	douyinRouter.POST("/user/register/", controller.UserRegister)
	douyinRouter.POST("/user/login/", controller.UserLogin)

	addr := strconv.Itoa(Config.GetInt("server.port"))
	if err := r.Run(":" + addr); err != nil {
		panic(err)
	}
}
