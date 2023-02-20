package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/favorite/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/favorite"
	"github.com/lius0712/douyin_server/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.VideoId)
	}
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errno.ErrBind
}
