package router

import (
	"github.com/gin-gonic/gin"
	. "saikumo.org/simple-douyin/src/config"
	"saikumo.org/simple-douyin/src/controller"
	"strconv"
)

func InitRouter() {
	r := gin.Default()

	douyinRouter := r.Group("/douyin")

	douyinRouter.POST("/user/register/", controller.UserRegister)

	addr := strconv.Itoa(Config.GetInt("server.port"))
	if err := r.Run(":" + addr); err != nil {
		panic(err)
	}
}
