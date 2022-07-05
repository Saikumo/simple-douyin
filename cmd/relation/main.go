package main

import (
	"log"
	relation "simple-douyin/cmd/relation/kitex_gen/relation/relationsrv"
)

func main() {
	svr := relation.NewServer(new(RelationSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}