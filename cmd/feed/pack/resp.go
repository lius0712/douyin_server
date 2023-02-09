package pack

import (
	"errors"
	"github.com/lius0712/douyin_server/kitex_gen/feed"
	"github.com/lius0712/douyin_server/pkg/errno"
)

func BuildVideoResp(err error) *feed.FeedResponse {
	if err == nil {
		return videoResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return videoResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return videoResp(s)
}

func videoResp(err errno.ErrNo) *feed.FeedResponse {
	return &feed.FeedResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  &err.ErrMsg,
	}
}
