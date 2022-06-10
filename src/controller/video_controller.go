package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/service"
	"strconv"
	"time"
)

func VideoFeed(c *gin.Context) {
	var latestTime int64
	var err error
	latestTimeStr := c.Query("latest_time")
	//如果时间戳为空则为当前时间
	if latestTimeStr == "" {
		latestTime = time.Now().Unix()
	} else {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			common.ResponseError(c, err)
			return
		}
	}

	videoFeedResponse, err := service.VideoFeed(latestTime)

	if err != nil {
		common.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, videoFeedResponse)
}

func videoFeedError(c *gin.Context, err error) {

	return
}
