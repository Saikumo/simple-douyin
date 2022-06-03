package main

import (
	. "saikumo.org/simple-douyin/src/config"
	"saikumo.org/simple-douyin/src/router"
	"strconv"
)

func main() {
	r := router.InitRouter()
	addr := strconv.Itoa(Config.GetInt("server.port"))

	if err := r.Run(":" + addr); err != nil {
		panic(err)
	}
}
