package main

import (
	"log"
	publish "saikumo.org/simple-douyin/kitex_gen/publish/publishsrv"
)

func main() {
	svr := publish.NewServer(new(PublishSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}