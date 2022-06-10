package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/dto"
)

func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, dto.Response{
		StatusCode: FAILURE_STATUS_CODE,
		StatusMsg:  err.Error(),
	})
}
