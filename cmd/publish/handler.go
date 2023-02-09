package main

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/publish/pack"
	"github.com/lius0712/douyin_server/cmd/publish/service"
	publish "github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishActionResp(errno.ErrorTokenInvalid)
		return resp, nil
	}
	if len(req.File) == 0 || len(req.Title) == 0 {
		resp = pack.BuildPublishActionResp(errno.ErrHttpBind)
		return resp, nil
	}
	err = service.NewPublishActionService(ctx).PublishAction(req, int(claim.Id))
	if err != nil {
		resp = pack.BuildPublishActionResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishActionResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}
	if req.UserId == 0 { //0默认为自己
		req.UserId = claim.Id
	}
	videos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
