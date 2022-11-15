package main

import (
	"context"
	publish "github.com/lius0712/douyin_server/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	//if len(req.File) == 0 || len(req.Title) == 0 {
	//	resp = pack.BuildPublishActionResp(errno.ErrHttpBind)
	//	return resp, nil
	//}
	//err := service.NewPublishActionService(ctx).PublishAction(req, )
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}
