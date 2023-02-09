package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/feed/pack"
	"github.com/lius0712/douyin_server/cmd/publish/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{
		ctx: ctx,
	}
}

func (s *PublishListService) PublishList(req *publish.PublishListRequest) (videos []*feed.Video, err error) {
	authorVideos, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(authorVideos) == 0 {
		return nil, nil
	}
	videos, err = pack.Videos(s.ctx, authorVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
