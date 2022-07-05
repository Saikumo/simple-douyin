package main

import (
	"log"
	feed "saikumo.org/simple-douyin/cmd/feed/kitex_gen/feed/feedsrv"
)

func main() {
	svr := feed.NewServer(new(FeedSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
