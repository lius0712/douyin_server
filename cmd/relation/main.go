package main

import (
	relation2 "github.com/lius0712/douyin_server/cmd/relation"
	relation "github.com/lius0712/douyin_server/kitex_gen/relation/relationservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(relation2.RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
