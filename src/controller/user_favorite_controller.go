package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/service"
	"strconv"
)

func UserFavorite(c *gin.Context) {
	userIdAny, _ := c.Get("user_id")
	userId := userIdAny.(uint)

	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseUint(videoIdStr, 10, 64)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseUint(actionTypeStr, 10, 64)
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	userFavoriteResponse, err := service.UserFavorite(userId, uint(videoId), uint(actionType))
	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, userFavoriteResponse)
}
