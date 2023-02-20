package pack

import (
	"context"
	"errors"
	"github.com/lius0712/douyin_server/cmd/comment/dal/db"
	dbUser "github.com/lius0712/douyin_server/cmd/user/dal/db"
	packUser "github.com/lius0712/douyin_server/cmd/user/pack"
	commentList "github.com/lius0712/douyin_server/kitex_gen/comment"
	"gorm.io/gorm"
)

func Comments(ctx context.Context, comment []*db.Comment, fromID int64) ([]*commentList.Comment, error) {
	comments := make([]*commentList.Comment, 0)
	for _, com := range comment {
		user, err := dbUser.QueryUserByID(ctx, int64(com.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		packUserInfo, err := packUser.User(ctx, user, fromID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &commentList.Comment{
			Id:         int64(com.ID),
			User:       packUserInfo,
			Content:    com.Content,
			CreateDate: com.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
