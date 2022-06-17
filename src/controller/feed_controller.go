package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/service"
)

func VideoFeed(c *gin.Context) {
	//纳秒 字符串
	latestTimeStr := c.Query("latest_time")

	videoFeedResponse, err := service.VideoFeed(latestTimeStr)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, videoFeedResponse)
}
