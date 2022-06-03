package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
}
