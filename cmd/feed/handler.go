package main

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/feed/pack"
	"github.com/lius0712/douyin_server/cmd/feed/service"
	feed "github.com/lius0712/douyin_server/kitex_gen/feed"
	"github.com/lius0712/douyin_server/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetUserFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	var uid int64 = 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrorTokenInvalid)
			return resp, err
		} else {
			uid = claim.Id
		}
	}
	videoList, nextTime, err := service.NewGetUserFeedService(ctx).GetUserFeed(req, uid)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}
	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = videoList
	resp.NextTime = &nextTime
	return resp, nil
}
