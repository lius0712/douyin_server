package pack

import (
	"errors"
	"github.com/lius0712/douyin_server/kitex_gen/publish"
	"github.com/lius0712/douyin_server/pkg/errno"
)

func BuildPublishActionResp(err error) *publish.PublishActionResponse {
	if err == nil {
		return publishActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return publishActionResp(s)
}

func publishActionResp(err errno.ErrNo) *publish.PublishActionResponse {
	return &publish.PublishActionResponse{
		Status: int32(err.ErrCode),
		Msg:    &err.ErrMsg,
	}
}

func BuildPublishListResp(err error) *publish.PublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return publishListResp(s)
}

func publishListResp(err errno.ErrNo) *publish.PublishListResponse {
	return &publish.PublishListResponse{
		Status: int32(err.ErrCode),
		Msg:    &err.ErrMsg,
	}
}
