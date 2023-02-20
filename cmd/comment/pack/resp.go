package pack

import (
	"errors"
	"github.com/lius0712/douyin_server/kitex_gen/comment"
	"github.com/lius0712/douyin_server/pkg/errno"
)

func BuildCommentActionResp(err error) *comment.CommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return commentActionResp(s)
}
func commentActionResp(err errno.ErrNo) *comment.CommentActionResponse {
	return &comment.CommentActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
func BuildCommentListResp(err error) *comment.CommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return commentListResp(s)
}
func commentListResp(err errno.ErrNo) *comment.CommentListResponse {
	return &comment.CommentListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
