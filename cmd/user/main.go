package main

import (
	"log"
	user "saikumo.org/simple-douyin/cmd/user/kitex_gen/user/usersrv"
)

func main() {
	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
