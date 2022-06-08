package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/service"
)

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userRegisterResponse, err := service.UserRegister(username, password)

	if err != nil {
		c.JSON(http.StatusOK, dto.UserRegisterResponse{
			Response: &dto.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, userRegisterResponse)
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userLoginResponse, err := service.UserLogin(username, password)

	if err != nil {
		c.JSON(http.StatusOK, dto.UserLoginResponse{
			Response: &dto.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, userLoginResponse)
}

func UserInfo(c *gin.Context) {
	userIdStr, _ := c.Get("user_id")
	userId := userIdStr.(int64)

	userInfoResponse, err := service.UserInfo(userId)

	if err != nil {
		c.JSON(http.StatusOK, dto.UserInfoResponse{
			Response: &dto.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, userInfoResponse)
}
