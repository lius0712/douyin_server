package main

import (
	feed2 "github.com/lius0712/douyin_server/cmd/feed"
	feed "github.com/lius0712/douyin_server/kitex_gen/feed/feedservice"
	"log"
)

func main() {
	svr := feed.NewServer(new(feed2.FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
