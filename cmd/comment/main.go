package main

import (
	"log"
	comment "simple-douyin/cmd/comment/kitex_gen/comment/commentsrv"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
