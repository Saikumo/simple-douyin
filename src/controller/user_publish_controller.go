package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/service"
	"strconv"
)

func UserPublish(c *gin.Context) {
	title := c.PostForm("title")
	userIdAny, _ := c.Get("user_id")
	userId := userIdAny.(uint)
	form, err := c.MultipartForm()
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	userPublishResponse, err := service.UserPublish(form, title, userId)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userPublishResponse)
}

func UserPublishList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	userPublishListResponse, err := service.UserPublishList(uint(userId))

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userPublishListResponse)
}
