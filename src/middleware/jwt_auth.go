package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/dto"
	"saikumo.org/simple-douyin/src/util"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从参数获取token
		token := c.Query("token")
		//拿不到token 再从body获取token
		if token == "" {
			token = c.PostForm("token")
		}
		//没有token
		if token == "" {
			c.JSON(http.StatusOK, dto.Response{
				StatusCode: 1,
				StatusMsg:  common.TokenIsNotExist.Error(),
			})
		}
		//校验并解析token
		claims, err := util.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusOK, dto.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}

		c.Set("user_id", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
}
