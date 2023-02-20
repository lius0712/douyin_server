package main

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/favorite/pack"
	"github.com/lius0712/douyin_server/cmd/favorite/service"
	favorite "github.com/lius0712/douyin_server/kitex_gen/favorite"
	"github.com/lius0712/douyin_server/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(errno.ErrorTokenInvalid)
		return resp, nil
	}
	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}
	if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
		resp = pack.BuildFavoriteActionResp(errno.ErrBind)
		return resp, nil
	}
	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(err)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteListResp(errno.ErrorTokenInvalid)
		return resp, nil
	}
	if req.UserId == 0 || claim.Id == 0 {
		resp = pack.BuildFavoriteListResp(errno.ErrBind)
		return resp, nil
	}
	videos, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
