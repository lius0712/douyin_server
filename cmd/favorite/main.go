package main

import (
	favorite2 "github.com/lius0712/douyin_server/cmd/favorite"
	favorite "github.com/lius0712/douyin_server/kitex_gen/favorite/favoriteservice"
	"log"
)

func main() {
	svr := favorite.NewServer(new(favorite2.FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
