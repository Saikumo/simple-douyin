package main

import (
	"fmt"
	. "saikumo.org/simple-douyin/src/config"
	"strconv"
)

func main() {
	addr := strconv.Itoa(Config.GetInt("server.port"))
	fmt.Println(addr)
}
