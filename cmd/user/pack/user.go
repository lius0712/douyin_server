package pack

import (
	"context"
	"github.com/lius0712/douyin_server/cmd/user/dal/db"
	"github.com/lius0712/douyin_server/kitex_gen/user"
)

func User(ctx context.Context, u *db.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{Name: "已注销用户"}, nil
	}
	followCount := int64(u.FollowCount)
	followerCount := int64(u.FollowerCount)
	isFollow := false
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}
