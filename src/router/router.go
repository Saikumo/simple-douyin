package router

import (
	"github.com/gin-gonic/gin"
	"saikumo.org/simple-douyin/src/controller"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	douyinRouter := r.Group("/douyin")

	douyinRouter.POST("/user/register", controller.UserRegister)

	return r
}
