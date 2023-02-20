package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/favorite/dal/db"
	"github.com/lius0712/douyin_server/cmd/favorite/pack"
	"github.com/lius0712/douyin_server/kitex_gen/favorite"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*feed.Video, error) {
	FavoriteVideos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	videos, err := pack.FavoriteVideos(s.ctx, FavoriteVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
