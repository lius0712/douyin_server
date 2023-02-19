package db

import (
	"context"
	"fmt"
	"github.com/lius0712/douyin_server/cmd/feed/dal/db"
	"github.com/lius0712/douyin_server/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string     `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string     `gorm:"type:varchar(256);not null" json:"password"`
	FavoriteVideos []db.Video `gorm:"many2many:t_favorite" json:"favorite_videos"`
	//Salt          string `gorm:"type:not null" json:"salt"`
	FollowCount   int64 `gorm:"default:0" json:"follow_count"`
	FollowerCount int64 `gorm:"default:0" json:"follower_count"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)
	fmt.Println("%%%%%%%%%%%")
	fmt.Println(userID)
	fmt.Println(res)
	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	fmt.Println(res)
	fmt.Println("nil nil nil")
	return res, nil
}
