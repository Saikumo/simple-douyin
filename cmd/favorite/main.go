package main

import (
	"log"
	favorite "saikumo.org/simple-douyin/kitex_gen/favorite/favoritesrv"
)

func main() {
	svr := favorite.NewServer(new(FavoriteSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
