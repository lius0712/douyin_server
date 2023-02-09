package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lius0712/douyin_server/cmd/api/rpc"
	"github.com/lius0712/douyin_server/cmd/feed/pack"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"github.com/lius0712/douyin_server/pkg/errno"
	"strconv"
)

func Feed(c *gin.Context) {
	var feedVar FeedParam
	var latestTime int64
	time := c.Query("latest_time")
	token := c.Query("token")
	if len(time) != 0 {
		if time, err := strconv.Atoi(time); err != nil {
			SendResponse(c, pack.BuildVideoResp(errno.DecodingFailed))
			return
		} else {
			latestTime = int64(time)
		}
	}
	feedVar.LatestTime = &latestTime
	feedVar.Token = &token
	resp, err := rpc.GetUserFeed(context.Background(), &feed.FeedRequest{
		LatestTime: feedVar.LatestTime,
		Token:      feedVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildVideoResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
