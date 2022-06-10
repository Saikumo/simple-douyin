package middleware

import (
	"github.com/gin-gonic/gin"
	"saikumo.org/simple-douyin/src/common"
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
			common.ResponseError(c, common.TokenIsNotExist)
			return
		}
		//校验并解析token
		claims, err := util.CheckToken(token)
		if err != nil {
			common.ResponseError(c, err)
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
}
