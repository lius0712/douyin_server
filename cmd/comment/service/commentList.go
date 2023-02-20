package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/comment/dal/db"
	"github.com/lius0712/douyin_server/cmd/comment/pack"
	"github.com/lius0712/douyin_server/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}
func (s *CommentListService) CommentList(req *comment.CommentListRequest, fromID int64) ([]*comment.Comment, error) {
	comments, err := db.GetVideoComments(s.ctx, int(req.VideoId))
	if err != nil {
		return nil, err
	}
	commentList, err := pack.Comments(s.ctx, comments, fromID)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}
