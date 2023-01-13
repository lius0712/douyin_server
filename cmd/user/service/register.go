package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/user"
	"github.com/lius0712/douyin_server/pkg/errno"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateService(ctx context.Context) *CreateUserService {
	return &CreateUserService{
		ctx: ctx,
	}
}

func (s *CreateUserService) CreateUser(req *user.UserRegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		Password: password,
	}})
}
