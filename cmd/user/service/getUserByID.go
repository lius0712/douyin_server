package service

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/cmd/user/pack"
	"github.com/lius0712/douyin_server/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{
		ctx: ctx,
	}
}

func (s *GetUserService) GetUser(req *user.GetUserRequest, fromID int64) (*user.User, error) {
	modelUser, err := db.QueryUserByID(s.ctx, req.Token.UserId)
	if err != nil {
		return nil, err
	}

	user, err := pack.User(s.ctx, modelUser, fromID)

	if err != nil {
		return nil, err
	}
	return user, nil
}
