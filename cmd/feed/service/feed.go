package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	"github.com/lius0712/douyin_server/cmd/feed/pack"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"time"
)

const LIMIT = 30 //返回视频最大数
type GetUserFeedService struct {
	ctx context.Context
}

func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{
		ctx: ctx,
	}
}
func (s *GetUserFeedService) GetUserFeed(req *feed.FeedRequest, fromId int64) (v []*feed.Video, nextTime int64, err error) {
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return v, nextTime, err
	}
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return v, nextTime, nil
	} else {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	if v, err = pack.Videos(s.ctx, videos, &fromId); err != nil {
		nextTime = time.Now().UnixMilli()
		return v, nextTime, err
	}
	return v, nextTime, nil
}