package main

import (
	"log"
	comment "saikumo.org/simple-douyin/kitex_gen/comment/commentsrv"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
