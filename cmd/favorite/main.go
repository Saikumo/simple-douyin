package main

import (
	"log"
	favorite "simple-douyin/cmd/favorite/kitex_gen/favorite/favoritesrv"
)

func main() {
	svr := favorite.NewServer(new(FavoriteSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
