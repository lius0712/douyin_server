package main

import (
	main2 "github.com/lius0712/douyin_server/cmd/publish"
	publish "github.com/lius0712/douyin_server/kitex_gen/publish/publishservice"
	"log"
)

func main() {
	svr := publish.NewServer(new(main2.PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
