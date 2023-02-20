package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/comment/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/comment"
	"github.com/lius0712/douyin_server/pkg/errno"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{
		ctx: ctx,
	}
}
func (s *CommentActionService) CommentAction(req *comment.CommentActionRequest) error {
	if req.ActionType == 1 {
		return db.AddComment(s.ctx, &db.Comment{
			UserID:  int(req.UserId),
			VideoID: int(req.VideoId),
			Content: *req.CommentText,
		})
	}
	if req.ActionType == 2 {
		return db.DeleteComment(s.ctx, *req.CommentId, req.VideoId)
	}
	return errno.ErrBind
}
