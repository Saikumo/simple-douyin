package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/service"
)

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userRegisterResponse, err := service.UserRegister(username, password)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userRegisterResponse)
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userLoginResponse, err := service.UserLogin(username, password)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userLoginResponse)
}

func UserInfo(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userId := userIdStr.(uint)

	userInfoResponse, err := service.UserInfo(userId)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userInfoResponse)
}
